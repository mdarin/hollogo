//
// Function panic and recovery
//
package main

import(
	"fmt"
	"os"
)

func main() {
	var fname string = "fupare.go"

	// Just panic :)
	// write(fname, nil)

	// Recover works in tandem with panic. A call to function 
	//recover returns the value that was passed as an argument to panic.


	// To be able to recover from an unwinding panic sequence, 
	// the code must make a deferred call to the recover function.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to make anagram: ", r)
		}
	}()
	//anagrams := "" 
	//write(fname, anagrams)
	write(fname, nil)
}


func write(fname string, anagrams map[string][]string) {
	_,err := os.OpenFile(fname, os.O_WRONLY+os.O_CREATE+os.O_EXCL, 0644)
	if err != nil {
		msg := fmt.Sprintf("Unable to create output file: %v", err)
		panic(msg)
	}
	//...
}


