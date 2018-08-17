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


// anonymous funcs
var (
	mul = func(op0, op1 int) int {
		return op0 * op1
	}

	sqr = func(val int) int {
		return mul(val,val)
	}
)


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


	fmt.Printf("Avg([1,2.5,3.75]) = %.2f\n", avg(1,2.5,3.75))
	points := []float64{9,4,3.7,7.1,7.9,9.2,10}
	// The slice can be passed as a variadic parameter 
	// by adding ellipses to the parameter in the sum(points...)
	//function call.
	fmt.Printf("Sum(%v) = %.2f\n", points, sum(points...))

	// an Euclidian division
	q,r := EuclidianDiv(71,5)
	fmt.Printf("EuclidianDiv(71,5): q = %d, r = %d\n", q, r)
	q,r = EuclidianDivNamed(142,7)
	fmt.Printf("EuclidianDivNamed(142,7): q = %d, r = %d\n", q, r)
	// The return keyword is followed by the number of result values matching (respectively)
	// the declared results in the function's signature. In the previous example, the signature of the
	// div function specifies two int values to be returned as result values. Internally, the
	// function defines int variables p and r that are returned as result values upon completion of
	// the function. Those returned values must match the types defined in the function's
	// signature or risk compilation errors.

	// Functions with multiple result values must be invoked in the proper context:
	// * They must be assigned to a list of identifiers of the same types respectively
	// * They can only be included in expressions that expect the same number of
	//   returned values


	// There is no inherent concept of passing parameter values by reference.
	// This means a local copy of the passed values is created inside the called function.
	someValue := math.Pi
	fmt.Printf("before dbl() : %.5f\n", someValue)
	dbl(someValue)
	fmt.Printf("after dbl() : %.5f\n", someValue)


	// Achieving pass-by-reference
	num := 2.807770
	fmt.Printf("num = %f\n", num)
	half(&num)
	fmt.Printf("half(num) = %f\n", num)


	//
	// Anonymous Functions and Closures
	//
	fmt.Printf("mul(25,7) = %d\n", mul(25,7))
	fmt.Printf("sqr(13) = %d\n", sqr(13))
	// Invoking anonymous function literals
	fmt.Printf(
		"94 (*C) = %.2f (*C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(94), // remember about last comma
	)

	for i := 0.0; i < float64(360); i += float64(45) {
		// the function literal code block, func() float64 {return deg
		// * math.Pi / 180}() , is defined as an expression that converts degrees to radians. With
		// each iteration of the loop, a closure is formed between the enclosed function literal and the
		// outer non-local variable, i . This provides a simpler idiom where the function naturally
		// accesses non-local values without resorting to other means such as pointers.
		rad := func() float64 {
			return i * math.Pi / 180
		}()
		fmt.Printf("%.2f Dec = %.2f Rad\n", i, rad)

		// attantion
		rad2 := func() float64 {
			return i * math.Pi / 180
		}
		fmt.Printf("address %.2f Dec = %.2f Rad\n", i, rad2)
		fmt.Printf("value %.2f Dec = %.2f Rad\n", i, rad2())
	}

	//
	//	Higher-order functions
	//

	// We have already established that Go functions are values bound to a type. So, it should not
	// be a surprise that a Go function can take another function as a parameter and also return a
	// function as a result value. This describes the notion known as a higher-order function,
	// which is a concept adopted from mathematics. While types such as struct let
	// programmers abstract data, higher-order functions provide a mechanism to encapsulate
	// and abstract behaviors that can be composed together to form more complex behaviors.
	nums := []int{4,32,11,67,2346,234,56,24,67}
	result := apply(nums, func(i int) int {
		return i / 2
	})
	result()
	// As you explore this book, and the Go language, you will continue to encounter usage of
	// higher-order functions. It is a popular idiom that is used heavily in the standard libraries.
	// You will also find higher-order functions used in some concurrency patterns to distribute workloads
}

func apply(nums []int, f func (int) int) func() {
	for i, v := range nums {
		nums[i] = f(v)
	}
	return func() {
		fmt.Println("Apply return: ", nums)
	}
}

func half(val *float64) {
	fmt.Printf("call half %f)\n", *val)
	*val = *val / 2
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

//
// Variadic parameters
//
// The last parameter of a function can be declared as variadic (variable length arguments) by
// affixing ellipses ( ... ) before the parameter's type. This indicates that zero or more values of
// that type may be passed to the function when it is called.
func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t += v
	}
	return t / float64(n)
}

func sum(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}


// a function that implements an Euclidian division algorithm
func EuclidianDiv(op0, op1 int) (int,int) {
	r := op0
	q := 0
	for r >= op1 {
		q++
		r = r - op1
	}
	return q,r
}


func EuclidianDivNamed(dvdn, dvsr int) (q,r int) {
	r = dvdn
	for r >= dvsr {
		q++
		r = r - dvsr
	}
	// Notice the return statement is naked;
	return
}

func dbl(val float64) {
	val = 2 * val // update param
	fmt.Printf("dbl() = %.5f\n", val)
}


