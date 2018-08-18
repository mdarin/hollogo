//
// Error signaling and handling
//
package main

import (
	"fmt"
	"errors"
	"os"
)

// The basic idea 
//
// The if...not...nil error handling idiom may seem excessive and
// verbose to some, especially if you are coming from a language with formal exception
// mechanisms. However, the gain here is that the program can construct a robust execution
// flow where programmers always know where errors may come from and handle them
// appropriately.

// Go has a simplified approach to error signaling and error
// handling that puts the onus on the programmer to handle 
// possible errors immediately aftera called function returns.

func main() {
	// In Go, the traditional way of signaling  errors is to 
	// return a value of type error when something goes wrong 
	// during the execution of your function.

	// illustration how to signal and handle an erorr	

	// simple error
	_, err := someFunc()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	// paramtrized error
	_, errParam := someAnotherFunc()
	if errParam != nil {
		fmt.Println("Error: ", errParam)
		// handle error
		os.Exit(1)
	}

}

func someAnotherFunc() (string,error) {
	if true {
		return "", fmt.Errorf("Parmetrized error '%s: %s'", "Kernel", "Panic")
	}
	return "nothing", nil
}

func someFunc() (string, error) {
	// if there is no error yet then... 
	if true {
		// let's imagin that something went wrong
		return "", errors.New("An orbitrary error occured :) Keep calm an stay happy.")
	}
	return "This point never be achived ;)", nil
}
