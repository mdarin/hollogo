//
// Parallelism in Go
// The sync package
//
package main

import(
	"fmt"
	"sync"
)

const WORKERS = 4

func main() {
 	values := make(chan int)
  result := make(chan int, WORKERS)
  var workersGroup sync.WaitGroup

  // generator
  go func() {
    for i := 0; i < 4; i++ {
      values<- i
    }
    close(values)
  }()

  // worker
  workerTask := func() {
    defer workersGroup.Done() // signal when done
    sum := 0
    for i := range values {
      sum += i
    }
    result<- sum
  }

  workersGroup.Add(WORKERS)

  // launch workers
  for i := 0; i < WORKERS; i++ {
    go workerTask()
  }

  workersGroup.Wait() // wait for all

  close(result)

  total := 0
  // gather partial results
  for partial := range result {
    total += partial
  }
  fmt.Println("total:", total)

} // eof main

