//
// Concurrency is considered to be the one of the most attractive features of Go.
//
// very unstable behaviour!
//
package main

import(
	"fmt"
	"strings"
	"time"
	"sync" // The sync package
)
// -----------------------------------------------------------------------------------------
// Go has its own concurrency primitive called the goroutine, which allows a program to
// launch a function (routine) to execute independently from its calling function. Goroutines
// are lightweight execution contexts that are multiplexed among a small number of OS-
// backed threads and scheduled by Go's runtime scheduler. That makes them cheap to create
// without the overhead requirements of true kernel threads. As such, a Go program can
// initiate thousands (even hundreds of thousands) of goroutines with minimal impact on
// performance and resource degradation.
// ------------------------------------------------------------------------------------------


// Channels
// ----------
// This is one area where Go diverges from its C lineage. Instead of having concurrent code
// communicate by using shared memory locations, Go uses channels as a conduit between
// running goroutines to communicate and share data.
// (*Do not communicate by sharing memory; instead, share memory by communicating.*)

// The channel type declares a conduit within which only values of a given element type may
// be sent or received by the channel.
// 
// chan <element type>
var channelNumric chan int

// the Go language supports the select statement that multiplexes selection among 
// multiple send and receive operations:
// select {
//	case <send_ or_receive_expression>:
//	default:
// }


// One popular idiom that is commonly encountered with Go concurrency is the use of the
// select statement, introduced previously, to implement timeouts. This works by using the
// select statement to wait for a channel operation to succeed within a given time duration
// using the API from the time package

type Service struct {
	started bool
	stpCh chan struct{}
	mutex sync.Mutex
}

func (s *Service) Start() {
	s.stpCh = make(chan struct{})
	go func() {
		s.mutex.Lock()
		s.started = true
		s.mutex.Unlock()
		fmt.Println(" mutex: Lock")
		<-s.stpCh // wait to be closed
	}()
}

func (s *Service) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.started {
		s.started = false
		close(s.stpCh)
	}
	fmt.Println(" mutex: Unlock")
}


type Service2 struct {
	started bool
	stpCh chan struct{}
	mutex sync.RWMutex
	cache map[int]string
}

func (s *Service2) Start2() {
	s.stpCh = make(chan struct{})
	s.cache = make(map[int]string)
	go func() {
		s.mutex.Lock()
		s.started = true
		s.cache[1] = "HAL-9000"
		s.cache[2] = "T-800"
		s.cache[3] = "T-1000"
		s.cache[4] = "R2D2"
		s.cache[5] = "WALL-E"
		for k, v := range s.cache {
			fmt.Println(" cache ", k, " ", v)
		}
		s.mutex.Unlock()
		<-s.stpCh // wait to be closed
	}()
}

func (s *Service2) Serve(id int) {
	s.mutex.RLock()
	msg := s.cache[id]
	s.mutex.RUnlock()
	if msg != "" {
		fmt.Println("serve: ", msg)
	} else {
		fmt.Println("serve: nothing, hasta la vista baby!")
	}
}


const MAX = 1000
const WORKERS = 4 

func main() {
	fmt.Println("hello concurrent world!")

	// The go statement
	// go <function or expression>
	// let's start number of workers
	go count(1, 10, 50, 10)
	go count(2, 60, 100, 10)
	go count(3, 110, 200, 20)
	// Goroutines may also be defined as function literals directly in the go statement
	go func() {
		count(4, 40, 60, 10)
	}()
	// The function literal provides a convenient idiom that allows programmers 
	// to assemble logic directly at the site of the go statement
	id := 5
	start := 30
	stop := 60
	step := 10
	go func() {
		count(id, start, stop, step)
	}()

	starts := []int{10,40,70,100}
	for _, j := range starts {
		go func(s int) {
			count(s, s+20, s+50, 10)
		}(j)
	}

	// When the make function is invoked without the capacity argument, 
	// it returns a bidirectional unbuffered channel.
	chUnbuf := make(chan int) // unguffered channel
	//chUnbuf <- 12 // fatal error: all goroutines are asleep - deadlock!
	go func() { chUnbuf <- 12 }()
	fmt.Println("unbuffered channel value: ", <-chUnbuf)

	// When the make function uses the capacity argument, 
	// it returns a bidirectional buffered channel
	chBuf := make(chan int, 4) // buffered channel
	chBuf <- 2
	chBuf <- 4
	chBuf <- 6
	chBuf <- 8
	// бадыль какой-то, но как-то так и делают, похоже...
	for k := 0; k < cap(chBuf); k++ {
		fmt.Println("buffered channel", k, " value: ", <-chBuf)
	}

	chClosed := make(chan int, 4)
	chClosed <- 2
	chClosed <- 4
	close(chClosed)
	for i := 0; i < 4; i++ {
		if val, opened := <-chClosed; opened {
			fmt.Println("channel value: ", val)
		} else {
			fmt.Println("channel closed ", val)
		}
	}

	// For now, let us use fmt.Scanln() to block and wait for keyboard input, as
	// shown in the following sample. In this version, the concurrent functions get a chance to
	// complete while waiting for keyboard input
	fmt.Scanln() // block for keyb input


	// ===========================
	// Writing concurrent programs
	// ===========================

	// a word histogram. 
	// The program reads the words from the data slice then, on a separate goroutine,
	// collects the occurrence of each word
	data := []string{
		"The yellow fish swims slowly in the water",
		"The brown dog barks loudly after a drink ...",
		"The dark bird bird of prey lands on a small ...",
	}

	histogram := make(map[string]int)
	done := make(chan bool)

	// anonymous function splits and count words
	go func() {
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				histogram[word]++
			}
		}
		done <- true
	}()

	// receive done state from channel and test it
	if state := <-done; state {
		for k, v := range histogram {
			fmt.Printf(" %s\t(%d)\n", k, v)
		}
	}

	// Streaming data
	// --------------
	// A natural use of channels is to stream data from one goroutine to another. This pattern is
	// quite common in Go code and for it to work, the followings must be done:
	// - Continuously send data on a channel
	// - Continuously receive the incoming data from that channel
	// - Signal the end of the stream so the receiver may stop

	histogram2 := make(map[string]int)
	wordsCh := make(chan string)

	// splits lines and sends words to channel
	go func() {
		defer close(wordsCh) // close channel when done
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.ToLower(word)
				wordsCh <-word
			}
		}
	}()


	// Using for...range to receive data
	// The pattern is so common in Go that the idiom is built into the language in the
	// form of the following for...range statement:
	// for <elemem> := range <channel>{...}
	// process word stream and count words
	// loop until wordsCh is closed
	//for {
	//	word, opened := <-
	//	wordsCh
	//	if !opened {
	//		break
	//	}
	//	histogram[word]++
	//}
	for word := range wordsCh {
		histogram2[word]++
	}
	// When the channel is closed (from the goroutine), the loop automatically breaks.

	fmt.Println("---------------")
	for k, v := range histogram2 {
		fmt.Printf("%s\t(%d)\n", k, v)
	}

	//
	// Generator functions
	//
	histogram3 := make(map[string]int)
	// In this example, the generator function, declared as func words(data []string) <-
	// chan string , returns a receive-only channel of string elements. The consumer function, in
	// this instance main() , receives the data emitted by the generator function, which is
	// processed using a for...range loop
	words3 := words(data) // returns handle to data channel
	for word := range words3 {
		histogram3[word]++
	}

	fmt.Println("------3---------")
	for k, v := range histogram3 {
		fmt.Printf("%s\t(%d)\n", k, v)
	}


	//
	// Selecting from multiple channels
	//
	histogram4 := make(map[string]int)
	stopCh := make(chan struct{}) // used to signal 'stop'

	words4 := words4(stopCh, data) // returns handle to channel
	for word := range words4 {
		if histogram4["the"] == 2 {
			close(stopCh)
		}
		histogram4[word]++
	}
	fmt.Println("-------4--------")
	for k, v := range histogram4{
		fmt.Printf("%s\t(%d)\n", k, v)
	}


	//
	// Channel timeout
	//
	histogram5 := make(map[string]int)
	done5 := make(chan struct {})

	go func() {
		defer close(done5)
		words5 := words(data) // returns handle to data channel
		for word := range words5 {
			histogram5[word]++
		}
		fmt.Println("-------5--------")
		for k, v := range histogram5 {
			fmt.Printf("%s\t(%d)\n", k, v)
		}
	}()

	select {
	case <-done:
		fmt.Println("Done counting words!")
	case <-time.After(200 * time.Microsecond):
		fmt.Println("Sorry, took too long to count.")
	}

	//
	// Synchronizing with mutex locks
	//
	fmt.Println(" =================== using mutex...")
	s := &Service{}
	s.Start()
	time.Sleep(time.Second) // imitation of doing some kind of work
	s.Stop()
	fmt.Println(" =================== done!")

	// The sync package also offers the RWMutex (read-write mutex), which can be used in cases
	// where there is one writer that updates the shared resource, while there may be multiple
	// readers.

	s2 := &Service2{}
	s2.Start2()
	s2.Serve(1)
	s2.Serve(2)
	s2.Serve(3)
	s2.Serve(4)
	s2.Serve(5)

	fmt.Println("----------------------------")
	// Concurrency barriers with sync.WaitGroup
	// Using WaitGroup requires three things:
	// - The number of participants in the group via the Add method
	// - Each goroutine calls the Done method to signal completion
	// - Use the Wait method to block until all goroutines are done
	values := make(chan int, MAX)
	result := make(chan int, 2)
	var wg sync.WaitGroup
	// create group about 2 workers	
	wg.Add(2)
	// gen multiple of 3 & 5 values
	go func() {
		for i := 1; i < MAX; i++ {
			if (1 % 3) == 0 || (i % 5) == 0 {
				values <-i // push downstream
			}
		}
		close(values)
	}()
	// work unit, calc partial result
	work := func() {
		defer wg.Done()
		r := 0
		for i := range values {
			r += i
		}
		result <-r
	}
	// distribute work to two goroutines
	go work()
	go work()
	// waiting for all workers until they done
	wg.Wait()
	// ======================================
	// It is important to remember that wg.Wait() will block indefinitely 
	// if its internal counter never reaches zero.
	// ======================================
	// gether partial result
	total := <-result + <-result
	fmt.Println("total: ", total)

	// Parallelism in Go
	// The Go runtime scheduler, by default, will create a
	// number of OS-backed threads for scheduling that is equal to the number of CPU cores. 
	// That quantity is identified by runtime value called GOMAXPROCS.

	values2 := make(chan int)
	result2 := make(chan int, WORKERS)
	var wg2 sync.WaitGroup

	go func() {
		for i := 1; i < MAX; i++ {
			if (i % 3) == 0 || (i % 5) == 0 {
				values2 <-i
			}
		}
		close(values2)
	}()

	work2 := func() {
		defer wg2.Done()
		r := 0
		for i := range values2 {
			r += i
		}
		result2 <-r
	}

	// lounch workers
	wg2.Add(WORKERS)
	for i := 0; i < WORKERS; i++ {
		go work2()
	}
	wg2.Wait()
	close(result2)
	total2 := 0
	// gether partial result
	for pr := range result2 {
		total2 += pr
	}
	fmt.Println("total2: ", total2)

} // eof main


func gen_mult_3_5(ch *chan int) {
	for i := 1; i < MAX; i++ {
		if (i % 3) == 0 || (i % 5) == 0 {
			*ch <-i
		}
	}
	close(*ch)
}


// selector
func words4(stopCh chan struct{}, data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out) // closes channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _, word := range words {
				word := strings.ToLower(word)
				select {
				case out <-word:
				case <-stopCh: // success first when close
					return
				} // eof select
			} // eof for words
		} // eof for line
	}()
	return out
}


// generator function that produces data
func words(data []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out) // close channel upon fn return
		for _, line := range data {
			words := strings.Split(line, " ")
			for _,word := range words {
				word = strings.ToLower(word)
				out <-word
			}
		}
	}()
	return out
}

// our worker
func count(id, start, stop, delta int) {
	fmt.Printf("worker[%d]\n", id)
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}



