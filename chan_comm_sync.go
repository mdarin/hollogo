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
const TIMEOUT = 1000

func main() {

 data := []string{
    "The yellow fish swims",
    "Blue is the water in a crystal glass",
    "Blue are the lilacs that grow in the grass",
    "Blue is delicious blueberry pie",
    "Blue are the sparkles in my cats eyes",
  }
	done := make(chan struct{})
	progress := make(chan string)
	wordsCounter := make(chan string)


	// process words worker
	go func() {
		defer close(done)
		fmt.Printf("\t\tReader\n\n")
		for word := range wordsGenerator(data) {
			fmt.Sprintf("%s", word)
			progress<- word
		}
	}()

	// progress monitor worker
	go func() {
		defer close(wordsCounter)
		fmt.Println("\t\tProgress\n")
		i := 0
		for word := range progress {
			i++
			fmt.Printf(".%d.",i)
			wordsCounter<- word
		}
	}()

	// counter worker
	go func() {
		defer close(progress)
		fmt.Printf("\t\tCounter\n\n")
		for word := range wordsCounter {
			fmt.Printf("%s", word)
	//	for range wordsCounter {
			continue
		}
	}()



	// controller-sycronizer
	select {
	case <-done:
		fmt.Println()
		fmt.Println()
		fmt.Println("Done counting words.\nTotal: ")
	case <-time.After(TIMEOUT * time.Microsecond):
//	case <-time.After(TIMEOUT * time.Millisecond):
		fmt.Println()
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

