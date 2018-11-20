//
// Methods, Interfaces, and Objects
// yet anather methods and interfaces exercise 
//
package main

import(
	"fmt"
	"math"
)

//
// Base type 
// that is narrowing a set of methods 
//
// It is important to note that the base type for method receivers cannot be a pointer (nor an interface).
//
type FuelTank float64
//
// Factory(constructor)
// Since Go does not support classes, there is no such concept as a constructor. However, one
// conventional idiom you will encounter in Go is the use of a factory function to create and
// initialize values for a type.
//
func NewFuelTank() *FuelTank {
	// declare and init the variable of FuelTank type
	//var ft FuelTank = 40
	// and use the address operator & (ampersand)
	//return &ft
	// The built-in function new(<type>) can also be used to initialize a pointer value. It first
	// allocates the appropriate memory for a zero-value of the specified type. The function then
	// returns the address for the newly created value.
	ft := new(FuelTank)
	*ft = 40
	return ft
}
//
// Method
// When a function is scoped to a type, or attached to the type, it is known as a method.
// A method is defined just like any other Go function. However, its definition includes a method receiver,
// which is an extra parameter placed before the method's name, used to specify the host type
// to which the method is attached.
//
func (ft *FuelTank) fill() {
	*ft = FuelTank(*ft * 4)
}




// methods...
type gallon float64

func newGallon(g gallon) *gallon {
	var gal *gallon = new(gallon)
	*gal = g
	return gal
}

func (g gallon) def(args ...gallon) {
	fmt.Println("n: ", len(args))
}

func (g gallon) quart() float64 {
	fmt.Println(" method:quart ::g -> ", g * 4)
	return float64(g * 4)
}

func (g gallon) half() {
	fmt.Println(" method:half ::g -> ", gallon(g * 0.5))
	g = gallon(g * 0.5)
}

func (g *gallon) double() {
	fmt.Println(" method:double ::g -> ", gallon(*g * 2))
	*g = gallon(*g * 2)
}



// interfaces
//
// Interfaces.
// The concept of interfaces in Go, similar to other languages, such as Java,
// is a set of methods that serves as a template to describe behavior.
// A Go interface, however, is a type specified by the interface{} literal,
// which is used to list a set of methods that satisfies the interface.
//
// Using idiomatic Go, an interface type is almost always declared as a named type .
//
//type shape interface {
//	area() float64
//	perim() float64
//}
//var s shape
//
// Implementing an interface
// The interesting aspect of interfaces in Go is how they are implemented and ultimately used.
//
// Implementing a Go interface is done implicitly. 
//
// There is no separate element or keyword required to indicate the intent of implementation. 
//
// NOTE: Any type that defines the method set of an interface type automatically satisfies its implementation.
//
type rect struct {
	name string
	length, height float64
}

// implementation of area() method of the shape interface
func (r *rect) area() float64 {
	return r.length * r.height
}
// implementation of perim() method of the shape interface
func (r *rect) perim() float64 {
	return 2 * r.length + 2 * r.height
}
//
//Subtyping with Go interfaces
//
type triangle struct {
	name string
	a, b, c float64
}
//
func (t *triangle) area() float64 {
	return 0.5*(t.a * t.b)
}
//
func (t *triangle) perim() float64 {
	return t.a + t.b + math.Sqrt((t.a*t.a) + (t.b*t.b))
}
//
func (t *triangle) String() string {
	format := "%s[sides: a=%.2f b=%.2f c=%.2f]"
	return fmt.Sprintf(format, t.name, t.a, t.b, t.c)
}
//
// Implementing multiple interfaces
// The implicit mechanism of interfaces allows any named type to satisfy multiple interface
// types at once. This is achieved simply by having the method set of a given type intersect
// with the methods of each interface type to be implemented.
//
// Interface embedding
// Another interesting aspects of the interface type is its support for type embedding
// (similar to the struct type). This gives you the flexibility to structure your types in ways
// that maximize type reuse.
// 
type shape interface {
	area() float64
}
//
type polygon interface {
	shape
	perim() float64
}
//
type curved interface {
	shape
	circonf() float64
}
//
type circle struct {
	name string
	rad float64
}
//
func (c* circle) area() float64 {
	return math.Pi * (c.rad * c.rad)
}
//
func (c *circle) circonf() float64 {
	return 2 * math.Pi * c.rad
}
//





// via interface v.1
// func shapeInfo(s shape) string {
// 	format := "Area = %.2f, Perim = %.2f"
// 	return fmt.Sprintf(format, s.area(), s.perim())
// }

// via polygon interface v.2
func polygonShapeInfo(p polygon) string {
	format := "Area = %.2f, Perim = %.2f"
	return fmt.Sprintf(format, p.area(), p.perim())
}

// via curved interface v.2
func curvedShapeInfo(c curved) string {
	format := "Area = %.2f, Circonf = %.2f"
	return fmt.Sprintf(format, c.area(), c.circonf())
}

//
// The empty interface type
//
// The interface{} type, or the empty interface type, is the literal representation of an
// interface type with an empty method set. According to our discussion so far, it can be
// deduced that all types implement the empty interface since all types can have a method set with
// zero or more members.
//
// NOTE: The empty interface is crucially important for idiomatic Go.
//
var emptyInterface interface{}


//
// Type assertion
//
// Type assertion is a mechanism that is available in Go to idiomatically narrow a variable 
// (of interface type) down to a concrete type and value that are stored in the variable.
//
type food interface {
	eat()
}

type veggie string
// vegied type implementation of the food interface
func (v veggie) eat() {
	fmt.Println("Eating", v)
}

type meat string
// meat type implementation of the food interface
func (m meat) eat() {
	fmt.Println("Eating tasty", m)
}

// hi-level API function with the food interface
func eat(f food) {
// The general form for type assertion expression is given as follows:
//
// <interface_variable>.(concrete type name)
//
// The type assertion expression can
// return two values: one is the concrete value (extracted from the interface) 
// and the second is a Boolean indicating the success of the assertion, as shown here:
//
// value, boolean := <interface_variable>.(concrete type name)
//
//	if v, ok := f.(veggie); ok {
//		switch v {
//		case "okra": fmt.Println("Yuk! not eating ", v)
//		default: v.eat()
//		} // eof switch
//	}
//
// A type assertion expression can also return just the value, as follows:
//
// value := <interface_variable>.(concrete type name)
//
// ATTENTION! This form of assertion is risky to do as the runtime will 
// cause a panic in the program if the value stored in the interface variable is not of 
// the asserted type. Use this form only if you have other safeguards to either prevent 
// or gracefully handle a panic.
	switch morsel := f.(type) {
	case veggie:
		switch morsel {
		case "okra":
			fmt.Println("Yuk! not eating", morsel)
		default:
			morsel.eat()
		}
	case meat:
		switch morsel {
		case
			"beef": fmt.Println("Yuk! not eating", morsel)
		default:
			morsel.eat()
		}
	default:
		fmt.Println("Not eating whatever that is", f)
	} // eof switch
}



//
// main driver
//
func main() {
	//var gal gallon = 5
	g := newGallon(5)
	g.def(1,2,3,4,4)
	fmt.Println("init: ", *g)
	g.quart()
	fmt.Println("quart: ", *g)
	g.half()
	fmt.Println("half: ", *g)
	g.double()
	fmt.Println("double: ", *g)

	r := &rect{
		name: "Red square",
		length: float64(4),
		height: float64(8),
	}
	//fmt.Println(t, "=>", shapeInfo(r))
	fmt.Println(r, "=>", polygonShapeInfo(r))

	t := &triangle{
		name: "Right triangle",
		a: float64(1),
		b: float64(2),
		c: float64(3),
	}
	//fmt.Println(t, "=>", shapeInfo(t))
	fmt.Println(t, "=>", polygonShapeInfo(t))

	c := &circle {
		name: "Found circle",
		rad: float64(6),
	}
	fmt.Println(c, "=>", curvedShapeInfo(c))

	var anyType interface{}
	anyType = float64(77) // I'm a float now"
	anyType = "I am s string now"
	fmt.Println("any type: ", anyType)

	printAnyType("The car is slow")
	var m map[string] string = map[string] string{"ID": "12345", "name": "Kerry"}
	printAnyType(m) // print out m as a map structure 
	printAnyType(1234567899) // print out an integer value


	eat(veggie("okra"))
	eat(veggie("cabage"))

	eat(meat("beef"))
	eat(meat("veal"))


} // eof main


func printAnyType(value interface{}) {
	fmt.Println("::value -> ", value)
}



