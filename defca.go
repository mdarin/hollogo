//
// Deferring function calls
//
package main

import (
	"fmt"
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
}
