//
// Syncronization via channels and communication 
// complex topology of workers
//
package main

import(
	"fmt"
	"strings"
	"time"
)

//
// You should restart number of times to see different rusult
//
// you may adjust this value to activate timeout case in the select statement
// it should be less then current value
// it's a boundary value now
const TIMEOUT = 8000

const WORKERS = 3

func main() {

 data := []string{
    "The yellow fish swims",
    "Blue is the water in a crystal glass",
    "Blue are the lilacs that grow in the grass",
    "Blue is delicious blueberry pie",
    "Blue are the sparkles in my cats eyes",
  }
	done := make(chan struct{})
	doneGroups := make(chan bool, 2)
	progress := make(chan string)
	wordsCounter := make(chan string)
	// words counting group 1
	counterAccumulator := make(chan int, WORKERS) // workers
	doneCounter := make(chan bool, WORKERS) // leader
	// words processing group 2
	processorAccumulator := make(chan string, WORKERS)
	doneProcessor := make(chan bool, WORKERS)


	// process words worker
	go func() {
		defer close(progress)
		fmt.Printf(" * Reader started")
		for word := range wordsGenerator(data) {
			progress<- word
		}
		fmt.Println()
		fmt.Println(" * Reader terminated")
	}()


	// progress monitor worker
	go func() {
		defer close(wordsCounter)
		fmt.Println(" * Progress sarted")
		i := 0
		for word := range progress {
			i++
			//fmt.Printf(".%d.",i)
			wordsCounter<- word
		}
		fmt.Println()
		fmt.Println(" * Progress termitated")
	}()

	// simple nested worker
	// counter worker
	go func() {
		fmt.Printf(" * Counter started")
		// child worker
		go func () {
			defer close(processorAccumulator)
			defer close(counterAccumulator)
			fmt.Println(" * Counter child")
			n := 0;
			for word := range wordsCounter {
				//fmt.Printf("{%s} ", word)
				n++
				// queue words, cycle for similar copy for every worker
			//	for i := 0; i < cap(processorAccumulator); i++ {
					processorAccumulator<- word
			//	}
			}
			// queue count result
			for marker := 0; marker < cap(counterAccumulator); marker++ {
				// markup for differ elements
				counterAccumulator<- n + marker
			}
			fmt.Println()
			fmt.Println(" * Counter child terminated")
		}()
		fmt.Println(" * Counter terminated")
	}()

	// ---------- GROUP-------------------
	// group or workers for processing queue
	// they dequeue values randomly
	for i := 0; i < cap(counterAccumulator); i++ {
		// Accountant worker
		go func(id int) {
			fmt.Printf(" * Accountant worker %d started\n", id)
			for wordsCount := range counterAccumulator {
				// worker's task
				fmt.Printf(" > Total %d: %d\n", id, wordsCount)
			}
			doneCounter<- true
			fmt.Println()
			fmt.Printf(" * Accountant worker %d terminated\n", id)
		}(i) // create worker ID
	}
	// group sycronizer(lider)
	go func() {
		fmt.Println(" * Group leader started")
		workersDoneCount := 0
		for range doneCounter {
			workersDoneCount++
			fmt.Printf(" > G1 workers %d done\n", workersDoneCount)
			if workersDoneCount >= WORKERS {
				// stop cycle and terminate group
				close(doneCounter)
				fmt.Println()
				fmt.Println(" ! Done counting words")
			}
		}
		// done goroup signale
		doneGroups<- true
		fmt.Println(" * Group leader terminated")
	}()
	// ----------END GROUP-------------------

	// nested comlex structure
	// starter parent worker
	go func() {
		// ---------- GROUP-------------------
		// group or workers for processing queue
		// they dequeue values randomly
		go func() {
			for i := 0; i < cap(processorAccumulator); i++ {
				// Word processor worker
				go func(id int) {
					//defer close(doneProcessor)
					fmt.Printf(" * Word processor worker %d started\n", id)
					for word := range processorAccumulator {
						// worker's task
						//fmt.Sprintf("%s",word)
						fmt.Printf(" > processor %d: %s\n", id, word)
					}
					// done when queue is empty
					doneProcessor<- true
					fmt.Println()
					fmt.Printf(" * Word processor worker %d terminated\n", id)
				}(i) // create worker ID
			}
		}()
		go func() {
		// group sycronizer(lider)
			go func() {
				fmt.Println(" * Group 2 leader started")
				workersDoneCount := 0
				for range doneProcessor {
					workersDoneCount++
					fmt.Printf(" > G2 workers %d done\n", workersDoneCount)
					if workersDoneCount >= WORKERS {
						// stop cycle and terminate group
						close(doneProcessor)
						// done goroup signale
						doneGroups<- true
						fmt.Println()
						fmt.Println(" ! Done processing words")
					}
				}
				fmt.Println(" * Group 2 leader terminated")
			}()
		}()
		// ----------END GROUP-------------------
	}()


	// All groups sycronizer(Supervisor)
	go func() {
		// signal everything done
		defer close(done)
		fmt.Println(" * Supervisor started")
		groupsDoneCount := 0
		for range doneGroups {
			groupsDoneCount++
			fmt.Printf(" # group %d done\n", groupsDoneCount)
			if groupsDoneCount >= 2 { // for all groups
				// stop cycle and terminate all supervised groups tree
				close(doneGroups)
				fmt.Println()
				fmt.Println(" ! Done supervised tree")
			}
		}
		fmt.Println(" * Supervisor terminated")
	}()

	// controller-sycronizer
	select {
	case <-done:
		fmt.Println()
		fmt.Println()
		fmt.Println(" ! Done")
	case <-time.After(TIMEOUT * time.Microsecond):
//	case <-time.After(TIMEOUT * time.Millisecond):
		fmt.Println()
		fmt.Println()
		fmt.Println(" ! Sorry, took too long to count")
	}

} // eof main


func wordsGenerator(data []string) <-chan string {
  outChan := make(chan string)
  go func() {
    defer close(outChan)
    for _, line := range data {
      words := strings.Split(line, " ")
      for _, word := range words {
        outChan<- word
      }
    }
  }()
  return outChan
}
