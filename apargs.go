//
// Accessing program arguments
//
package main

import(
	"fmt"
	"os"
)

func main() {
	// Position 0 in slice os.Args holds the fully qualified name of the program's binary path.
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}

