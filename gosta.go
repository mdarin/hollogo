// 
// go control flow
//
package main

import "fmt"


func main() {
	var num0 = 250

	if num0 > 100 || num0 < 900 {
		fmt.Println("in the interval")
	}

	if num0 > 300 || num0 < 900 {
		fmt.Println("out the interval")
	} else {
		fmt.Println("in the interval")
	}

	// The initialization statement follows normal variable declaration and initialization rules. The
	// scope of the initialized variables is bound to the if statement block, beyond which they
	// become unreachable. This is a commonly used idiom in Go
	if num1 := 319; num1 > 100 || num1 < 900 {
		fmt.Println("initialization -> ", num1)
	}
	var res = myFunc()
	fmt.Println("myfunc -> ", res)
}


//my firts function :) 
func myFunc() bool {
	return true
}
