//
// Syncronization via channels and communication 
//
package main

import(
	"fmt"
	"strings"
	"time"
)

// you may adjust this value to activate timeout case in the select statement
// it should be less then current value
// it's a boundary value now
const TIMEOUT = 10000

func main() {

 data := []string{
    "The yellow fish swims",
    "Blue is the water in a crystal glass",
    "Blue are the lilacs that grow in the grass",
    "Blue is delicious blueberry pie",
    "Blue are the sparkles in my cats eyes",
  }
	done := make(chan struct{})
	doneCounter := make(chan struct{})
	progress := make(chan string)
	wordsCounter := make(chan string)
	counterAccumulator := make(chan int)


	// process words worker
	go func() {
		//defer close(done)
		defer close(progress)
		fmt.Printf("*Reader started\n")
		for word := range wordsGenerator(data) {
			fmt.Sprintf("%s", word)
			progress<- word
		}
		fmt.Println("*Reader terminated\n")
	}()


	// progress monitor worker
	go func() {
		defer close(wordsCounter)
		fmt.Println("*Progress sarted\n")
		i := 0
		for word := range progress {
		//for range progress {
			i++
			fmt.Printf(".%d.",i)
			wordsCounter<- word
		}
		fmt.Println("*Progress termitated\n")
	}()

	// counter worker
	go func() {
		fmt.Printf("*Counter started\n")
		// child worker
		go func () {
			//defer close(doneCounter)
			defer close(counterAccumulator)
			fmt.Println("*Counter child\n")
			n := 0;
			for word := range wordsCounter {
				fmt.Printf("%s ", word)
		//	for range wordsCounter {
				n++
				continue
			}
			counterAccumulator<- n
			fmt.Println("*Counter child terminated\n")
		}()
		fmt.Println("*Counter terminated\n")
	}()

	// Accountant worker
	go func() {
		defer close(doneCounter)
		fmt.Println("*Result accumulator started")
		for wordsCount := range counterAccumulator {
			fmt.Println("Total:", wordsCount)
		}
		fmt.Println("*Result accumulator terminated")
	}()


	// controller-sycronizer
	select {
	case <-doneCounter:
		fmt.Println()
		fmt.Println("Done counting words.\n")
	case <-done:
		fmt.Println()
		fmt.Println("Done reading.")
	case <-time.After(TIMEOUT * time.Microsecond):
//	case <-time.After(TIMEOUT * time.Millisecond):
		fmt.Println()
		fmt.Println("Sorry, took too long to count.")
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

