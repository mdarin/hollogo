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
type shape interface {
	area() float64
	perim() float64
}
var s shape
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

// via interface
func shapeInfo(s shape) string {
	format := "Area = %.2f, Perim = %.2f"
	return fmt.Sprintf(format, s.area(), s.perim())
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
	fmt.Println(r, "=>", shapeInfo(r))

	t := &triangle{
		name: "Right triangle",
		a: float64(1),
		b: float64(2),
		c: float64(3),
	}
	fmt.Println(t, "=>", shapeInfo(t))

} // eof main



