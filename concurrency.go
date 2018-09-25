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

