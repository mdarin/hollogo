//
// Goroutines, channels and io.Readers/Writers
// more complex usage example
//
package main

import(
	"fmt"
	"os"
	//"io"
	"bufio"
//	"strings"
)

// serial application
func main() {
	// make originate srouce
	fin, err := os.Open("./test_case10.csv")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer fin.Close()

	// setup fin as a source origin for a scanner
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Scanner error:", err)
	}
}

