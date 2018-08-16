//
// Go Data types
//
package main

import(
	"fmt"
	"strings"
)

// pointers ^_^. 
// We're like pointers! Aren't we?
var valPtr *float32
var countPtr *int
var person *struct{
	name string
	age int
}
var matrix *[1024]int
var row *[]int64


func main() {
	// folat just for instance
	p := 3.1415926535
	e := .5772156649
	x := 7.2E-5
	y := 1.616199e-35
	z := .416833e32

	fmt.Println(p, e, x, y, z)

	// literals
	vals := []int{
		1024,
		0x0FF1CE,
		0x8BADF00D,
		0xBEEF,
		0777,
	}
	for _, v := range vals {
		if v == 0xBEEF {
			fmt.Printf("Go dec:%d hex: %X\n", v, v)
			break
		}
	}
	// complex
	cx := -3.5 + 2i
	fmt.Printf("complex: %v\n", cx)
	fmt.Printf("r: %+g, i: %+g\n", real(cx), imag(cx))
	var flag bool = true
	fmt.Println("Flag: ", flag)

	fmt.Println("Pointers: ", valPtr, countPtr, person, matrix, row)

	var a int = 1024
	var aptr *int = &a
	//var bptr *int = &2048 this is an compilation error!  a directly value pointer

	fmt.Printf("a = %v\n", a)
	fmt.Printf("aptr = %v an adress\n", aptr)
	fmt.Printf("*aptr = %v a poited value\n", *aptr)

	// structs' poiters
	structPtr := &struct{ x, y int} {44, 55}
	pairPtr := &[]string{"A", "B"}

	fmt.Printf("struct = %v, type = %T\n", structPtr, structPtr)
	fmt.Printf("*struct = %v, type = %T with start (*)\n", *structPtr, structPtr)
	fmt.Printf("pairPtr= %v, type = %T\n", pairPtr, pairPtr)
	fmt.Printf("pairPtr= %v, type = %T with start (*)\n", *pairPtr, pairPtr)

	// new()
	intptr := new(int)
	*intptr = 44

	pp := new(struct{ first,last string})
	pp.first = "Samuel"
	pp.last = "Pierre"

	fmt.Printf("Value: %d, type %T\n", *intptr, intptr)
	// now, I'm stay confusing on how to use pointers properly
	fmt.Printf("Person: %+v\n", pp)
	fmt.Printf("*Person: %+v\n", *pp)
	// but read further and learn more about how to do that.
	//
	// Pointer indirection – accessing referenced values
	//
	// If all you have is an address, you can access the value to which it points by 
	// applying the * operator to the pointer value itself (or dereferencing).
	aa := 33
	double(&aa)
	fmt.Println("Doubled: ", aa)
	capp(pp)
	fmt.Println("Capitalized: ", pp)

	// non-composite type decl
	type truth bool
	type quart float64
	type gallon float64
	type node string

	var c celsius = 32.0
	f := fahrenheit(122)
	fmt.Printf("Temperature: %.2f \u00b0C = %.2f \u00b0K\n", c, celToKel(c))
	fmt.Printf("Temperature: %.2f \u00b0F = %.2f \u00b0C\n", f, fharToCel(f))

	//
	// type conversin
	//
	type signal int
	// ===================
	// The expression actual + count causes a build time error because both variables are of
	// different types. Even though variables actual and count are of numeric types and int32
	// and int have the same memory representation, the compiler still rejects the expression.
	//var count int32
	//var actual int
	//var test int64 = actual + count
	// The same is true for declared named types and their underlying types. The compiler will
	// reject assignment var event int = sig because type signal is considered to be
	// different from type int . This is true even though signal uses int as its underlying type.
	//var sig signal
	//var event int = sig
	// ===================
	// ! Type conversion is done using the following format:
	// ! <target_type>(<value or expression>)
	var count int32
	var actual int
	var test int32 = int32(actual) + count
	var sig signal
	var event int = int(sig) 
	fmt.Println("converted test: ", test)
	fmt.Println("converted event: ", event)
}


func double(x *int) {
	*x = *x * 2 // *x *= 2 ?
}

func capp(p *struct{ first, last string }) {
	// However, when dealing with composites, the idiom is more forgiving. 
	// It is not necessary to write *p.first to access the
	// pointer's field value. We can drop the * and just use p.first = strings.ToUpper(p.first).
	p.first = strings.ToUpper(p.first)
	p.last = strings.ToUpper(p.last)
}



type fahrenheit float64
type celsius float64
type kelvin float64

func fharToCel(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9)
}

func fharToKel(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9 + 273.15)
}

func celToFahr(c celsius) fahrenheit {
	return fahrenheit(c * 5 / 9 + 32)
}

func celToKel(c celsius) kelvin {
	return kelvin(c + 273.15)
}

