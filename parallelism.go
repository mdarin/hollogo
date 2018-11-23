//
// Parallelism with sync package
//
package main

//
// map - reduce model
// worker as a map unit that produces data streams
// result as a reduce unit that processes data streams from multiple map units
//

import(
	"fmt"
	"os"
	"io"
	"sync"
)


const WORKERS = 4
const (
	STDOUT = 0
	FILE
)

func main() {
	// init consumers' channels
	consumers := []chan int{
		make(chan int, WORKERS),
		make(chan int, WORKERS),
	}

	// create and init group of workers 
	var workersGroup sync.WaitGroup
	workersGroup.Add(WORKERS)

	// source
	values := make(chan int)
	go func() {
		defer close(values)
		for i := 0; i < 4; i++ {
			values<- i
		}
	}()

	// worker
	workerTask := func(workerID int, consumers []chan int) {
		// signal when done
		defer workersGroup.Done()

		sum := 0
		var source chan int

		// define source
		if workerID % 2 == 0 {
			source = values
		} else {
			source = dataGenerator()
		}

		// read source and calc sum
		for i := range source {
			sum += i
		}

		// send sum to multiple consumers
		for _, consumer := range consumers {
			consumer<- sum
		}

	} // eof worker

	// launch goroup of workers
	for i := 0; i < WORKERS; i++ {
		go workerTask(i, consumers)
	}

	// wait for all workers done
	workersGroup.Wait()

	// close consumers' channels
	for _,consumer := range consumers {
		close(consumer)
	}

	// gather partial results
	total := 0
	for partial := range consumers[STDOUT] {
		total += partial
	}

	// stdout as a target
	fmt.Println("total:", total)

	// gather partial result2
	total = 0
	for partial := range consumers[FILE] {
		total += partial
	}

	// file as a target
	fout,_ := os.Create("./result.txt")
	defer fout.Close()
	io.WriteString(fout, fmt.Sprintf("%d\n",total))

} // eof main


// generator
func dataGenerator() chan int {
	values := make(chan int)
	go func() {
		defer close(values)
		for i := 100; i < 104; i++ {
			values<- i
		}
	}()
	return values
}
