//
// Deferring function calls
//
package main

import (
	"fmt"
	"os"
)

// Go supports the notion of deferring a function call. Placing the keyword defer before a
// function call has the interesting effect of pushing the function unto an internal stack,
// delaying its execution right before the enclosing function returns.

func do(steps ...string) {
	defer fmt.Println("* All done!")
	for _, s := range steps {
		defer fmt.Println(" -",s)
	}
	fmt.Println("* Starting")
}

func main() {
	var steps = []string{
		"Find key",
		"Aplly break",
		"Put key in ignition",
		"Start car",
	}

	do(steps...)

	//
	// Using defer
	//
	// The defer keyword modifies the execution flow of a program by delaying function calls.
	// One idiomatic usage for this feature is to do a resource cleanup.
	readFile("")
}

func readFile(fname string) ([]string, error) {
	//...
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//...
	return []string{""}, nil
	// The pattern of opening-defer-closing resources is widely used in Go. By placing the
	// deferred intent immediately after opening or creating a resource allows the code to read
	// naturally and reduces the likeliness of creating a resource leakage.
}
