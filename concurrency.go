//
// Concurrency is considered to be the one of the most attractive features of Go.
//
package main

import(
	"fmt"
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
}

// our worker
func count(id, start, stop, delta int) {
	fmt.Printf("worker[%d]\n", id)
	for i := start; i <= stop; i += delta {
		fmt.Println(i)
	}
}

