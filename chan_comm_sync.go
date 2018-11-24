//
// Syncronization via channels and communication 
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
const TIMEOUT = 800

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
	doneCounter := make(chan bool, WORKERS)
	progress := make(chan string)
	wordsCounter := make(chan string)
	counterAccumulator := make(chan int, WORKERS)


	// process words worker
	go func() {
		//defer close(done)
		defer close(progress)
		fmt.Printf(" * Reader started\n")
		for word := range wordsGenerator(data) {
			fmt.Sprintf("%s", word)
			progress<- word
		}
		fmt.Println(" * Reader terminated\n")
	}()


	// progress monitor worker
	go func() {
		defer close(wordsCounter)
		fmt.Println(" * Progress sarted\n")
		i := 0
		for word := range progress {
		//for range progress {
			i++
			fmt.Printf(".%d.",i)
			wordsCounter<- word
		}
		fmt.Println(" * Progress termitated\n")
	}()

	// counter worker
	go func() {
		fmt.Printf(" * Counter started\n")
		// child worker
		go func () {
			//defer close(doneCounter)
			//defer close(counterAccumulator)
			fmt.Println(" * Counter child\n")
			n := 0;
			for word := range wordsCounter {
				fmt.Printf("%s ", word)
		//	for range wordsCounter {
				n++
				continue
			}
			// queue result
			counterAccumulator<- n
			counterAccumulator<- n + 1 // markup for differ elements
			counterAccumulator<- n + 2 // markup for differ elements
			fmt.Println(" * Counter child terminated\n")
		}()
		fmt.Println(" * Counter terminated\n")
	}()

//FIXME: panic: close of closed channel
// and process panic 

	// group or workers for processing queue
	// they dequeue values randomly
	// Accountant1 worker
	go func() {
		defer close(doneCounter)
		fmt.Println(" * Result1 accumulator started\n")
		for wordsCount := range counterAccumulator {
			fmt.Println("Total1:", wordsCount)
			doneCounter<- true
		}
		fmt.Println(" * Result1 accumulator terminated\n")
	}()

	// Accountant2 worker
	go func() {
		defer close(doneCounter)
		fmt.Println(" * Result2 accumulator started\n")
		for wordsCount := range counterAccumulator {
			fmt.Println("Total2:", wordsCount)
			doneCounter<- true
		}
		fmt.Println(" * Result2 accumulator terminated\n")
	}()

	// Accountant3 worker
	go func() {
		defer close(doneCounter)
		fmt.Println(" * Result3 accumulator started\n")
		for wordsCount := range counterAccumulator {
			fmt.Println("Total3:", wordsCount)
			doneCounter<- true
		}
		fmt.Println(" * Result3 accumulator terminated\n")
	}()

	// group sycronizer(lider)
	go func() {
		defer close(done)
		fmt.Println(" * Group leader started\n")
		workersDoneCount := 0
		//for workerDone := range doneCounter {
		for range doneCounter {
			workersDoneCount++
			fmt.Println()
			fmt.Println("workers:", workersDoneCount)
			if workersDoneCount >= WORKERS {
				// termite group
				close(counterAccumulator)
				fmt.Println()
				fmt.Println("Done counting words\n")
			}
		}
		fmt.Println(" * Group leader terminated\n")
	}()


	// controller-sycronizer
	select {
	case <-done:
		fmt.Println()
		fmt.Println()
		fmt.Println("Done")
	case <-time.After(TIMEOUT * time.Microsecond):
//	case <-time.After(TIMEOUT * time.Millisecond):
		fmt.Println()
		fmt.Println()
		fmt.Println("Sorry, took too long to count")
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

