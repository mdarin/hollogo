//
// Functions in Go
// 
package main

import(
	"fmt"
	"math"
)

//
// Canonical form function declaration
//
// the 'func' keyword marks the beginning of a function declaration
// <func-id> an optional identifier that is used to name the function for future reference
// <arg-list> a required set of parentheses enclosing an optional list of fucntion arguments
// <result-list> an optional list of types for values returned by the function.
// when function returns more than one value, an enclosing set of parentheses is ruquired

// func [<func-id>] ([<arg-list>]) [(<result-list>)] {
// [return] [<value or expr-list>]
//}

// the 'return' statement causes the execution flow to exit function.
// when the function defines returned values in its definition, a return statement is 
// required as the last logical statement of the function. 
// Otherwise, if no return values arespecified, the return statement is optional.


func main() {
	printPi()
	fmt.Printf("Avogadro: %e 1/mol\n", avogadro())
	//fmt.Println("Anonymous: ", anoF())
	fib(41)
	var prime float32 = 37.78034
	// let's try a type convertion!
	fmt.Printf("isPrime(%d): %v\n", int(prime), isPrime(int(prime)))

	// it's look very intresting...
	var opAdd func(int, int) int = add
	opSub := sub
	// Can I define anonymous functin outside anther fucntion?
	// defined over here it is working good!
	anonymous := func () bool { return true }
	fmt.Println("anonymous: ", anonymous())
	fmt.Printf("opAdd(12,44) = %d\n", opAdd(12,44))
	fmt.Printf("opSub(99,13) = %d\n", opSub(99, 13))
}

//("fmt" "math") func printPi() {
//	fmt.Printf("printPi() %v\n", math.Pi)
//}

//func main() { printPi() }

func printPi() {
	fmt.Printf("printPi %v\n", math.Pi)
}

func avogadro() float64 {
	return 6.02214129e23
}


//var anoF func = func () bool { 
//	return true 
//}

// (!) Function signature
// The set of specified parameter types, result types, and the order in which
// those types are declared is known as the signature of the function. It is
// another unique characteristic that help identify a function. Two functions
// may have the same number of parameters and result values; however, if
// the order of those elements are different, then the functions have different
// signatures.

func fib(n int) {
	fmt.Printf("fib(%d):[", n)
	var p0,p1 uint64 = 0,1
	fmt.Printf(" %d %d ", p0,p1)
	for i := 2; i <= n; i++ {
		p0, p1 = p1, p0+p1
		fmt.Printf("%d ", p1)
	}
	fmt.Println("]")
}


func isPrime(n int) bool {
	// magic of type convertion...
	lim := int(math.Sqrt(float64(n)))
	for p := 2; p <= lim; p++ {
		if (n % p) == 0 {
			return false
		}
	}
	return true
}


func add(op0 int, op1 int) int {
	return op0 + op1
}

func sub(op0, op1 int) int {
	return op0 - op1
}


