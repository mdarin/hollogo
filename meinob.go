//
// Methods, Interfaces, and Objects
//
package main

import (
	"fmt"
	"math"
)



// Go methods
// A Go function can be defined with a scope narrowed 
// to that of a specific type. When a function is scoped 
// to a type, or attached to the type, it is known as a method.

// A method is defined just like any other Go function. 
// However, its definition includes a method receiver,
// which is an extra parameter placed before the method's name, 
// used to specify the host type to which the method is attached.

// It shows the quart method attached to the type -gallon- based
// receiver via the -g gallon- receiver parameter:
// receivers are normal function parameters.
type gallon float64

func (g gallon) quart() quart {
	return quart(g * 4)
}

type ounce float64
func (o ounce) cup() cup {
	return cup(o * 0.1250)
}

type cup float64
func (c cup) quart() quart {
	return quart(c * 0.25)
}

func (c cup) ounce() ounce {
	return ounce(c * 8.0)
}

type quart float64
func (q quart) gallon() gallon {
	return gallon(q * 0.25)
}

func (q quart) cup() cup {
	return cup(q * 4.0)
}

func (g gallon) half() {
	g = gallon(g * 0.5)
}
// It uses a method receiver of the *gallon type, 
// which is updated using *g = gallon(*g * 2)
func (g *gallon) halfPtr() {
	*g = gallon(*g * 0.5)
}

func (g *gallon) double() {
	*g = gallon(*g * 2)
}
// --------------------------------------------------
// Pointer receiver parameters are widely used in Go.
// --------------------------------------------------



// The struct as object
type fuel int
const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle struct {
	mark string
	model string
}

type engine struct {
	fuel fuel
	thrust int
}

func (e *engine) start() {
	fmt.Println(" Engine started.")
}

type truck struct {
	vehicle
	engine
	axels int
	wheels int
	class int
}

// The constructor function
// Since Go does not support classes, there is no such
// concept as a constructor. However, one conventional 
// idiom you will encounter in Go is the use of a factory 
// function to create and initialize values for a type.
func newTruck(mk, mdl string) *truck {
	return &truck {vehicle: vehicle{mk, mdl}}
}


func (t *truck) drive() {
	fmt.Printf(" Truck %s %s, on the go!\n", t.mark, t.model)
}

type plane struct {
	vehicle
	engine
	engineCount int
	fixedWings bool
	maxAltitude int
}

func newPlane(mk, mdl string) *plane {
	p := &plane{}
	p.mark = mk
	p.model = mdl
	return p
}

func (p *plane) fly() {
	fmt.Printf(
		" Aircraft %s %s clear for takeoff!\n",
		p.mark, p.model,
	)
}



//
// The interface type
// When you talk to people who have been doing Go for a while,
// they almost always list the interface as one of their favorite features of the language. 
// The concept of interfaces in Go, similar to other languages, such as Java, 
// is a set of methods that serves as a template to describe behavior. 
// A Go interface, however, is a type specified by the interface{} literal,
// which is used to list a set of methods that satisfies the interface.

// Using idiomatic Go, an interface type is almost always declared as a named type .

// Embedding becomes a crucial feature, especially when the code applies
// type validation using type checking. It allows a type to roll up type
// information, thus reducing unnecessary assertion steps
type shape interface {
	area() float64
}
var s shape

type polygon interface {
	shape
	perim() float64
}

type curved interface {
	shape
	circonf() float64
}

type rect struct {
	name string
	length float64
	height float64
}

func (r *rect) area() float64 {
	return r.length * r.height
}

func (r *rect) perim() float64 {
	return 2 * r.length + 2 * r.height
}

func (r *rect) String() string {
	return fmt.Sprintf(
		" %s[sides: l=%.2f h=%.2f]",
		r.name, r.length, r.height,
	)
}

type triangle struct {
	name string
	a float64
	b float64
	c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) perim() float64 {
	return t.a + t.b + math.Sqrt((t.a * t.a) + (t.b * t.b))
}

func (t *triangle) String() string {
	return fmt.Sprintf(
		" %s[sides: a=%.2f b=%.2f c=%.2f]",
		t.name, t.a, t.b, t.c,
	)
}

type circle struct {
	name string
	rad float64
}

func (c *circle) area() float64 {
	return math.Pi * (c.rad * c.rad)
}

func (c *circle) circonf() float64 {
	return 2 * math.Pi * c.rad
}
// capital letter 'S' in the String means this function is exported by this module 
func (c *circle) String() string {
	return fmt.Sprintf(
		" %s[radius: r=%.2f]",
		c.name, c.rad,
	)
}

func shapeInfo(s shape) string {
	return fmt.Sprintf(
		"Area = %.2f %.2f",
		s.area(),
	)
}

func polygonInfo(p polygon) string {
	return fmt.Sprintf(
		"Perim = %.2f",
		p.perim(),
	)
}

func curvedInfo(c curved) string {
	return fmt.Sprintf(
		"Circonf = %.2f",
		c.circonf(),
	)
}

// The empty interface type

// Method set
// ----------
// The number of methods attached to a type, via the receiver parameter, 
// is known as the type's method set.



// Type assertion
type food interface {
	eat()
}

type veggie string
func (v veggie) eat() {
	fmt.Println("Eating", v)
}

type meat string
func (m meat) eat() {
	fmt.Println("Eating tasty", m)
}


//
// main driver
//

func main() {
	g := gallon(5)
	// a method has the scope of a type. Therefore, it can only 
	// be accessed via a declared value (concrete or pointer) 
	// of the attached type using -dot notation-.
	fmt.Println(" quart: ", g.quart())
	ozs := g.quart().cup().ounce()
	fmt.Printf(" %.2f gallons = %.2f ounce\n", g, ozs)
	var gal gallon = 5
	gal.half()
	fmt.Println(" half: ", gal)
	g.halfPtr()
	fmt.Println(" halfptr: ", g)
	gal.double()
	fmt.Println(" double: ", gal)

	// create truck
	t := &truck{
		vehicle: vehicle{"Ford", "F750"},
		engine: engine{GASOLINE+BIO, 700},
		axels: 2,
		wheels: 6,
		class: 3,
	}
	t.start()
	t.drive()

	// create plane
	p := &plane{}
	p.mark = "HondaJet"
	p.model = "HA-420"
	p.fuel = JET
	p.thrust = 2050
	p.engineCount = 2
	p.fixedWings = true
	p.maxAltitude = 43000
	p.start()
	p.fly()

	r := &rect{"Square", 4.0, 4.0}
	// using Stirng virtual method to stringify rect struct
	fmt.Println(r, "=>", polygonInfo(r))
	tr := &triangle{"Right Triangle", 1,2,3}
	// using Stirng virtual method to stringify triangle struct
	fmt.Println(tr, "=>", polygonInfo(tr))
	cr := &circle{"Circle", 12.0}
	// using Stirng virtual method to stringify triangle struct
	fmt.Println(cr, "=>", curvedInfo(cr))


	// The empty interface is crucially important for idiomatic Go
	var anyType interface{}
	anyType = 77.0
	anyType = "I am a string now"
	fmt.Println(anyType)
	printAnyType("The car is slow")
	m := map[string] string{"ID":"1234567", "name": "Kerry"}
	printAnyType(m)
	printAnyType(124656763457)


	vg := veggie("okra")
	mt := meat("beef")
	vg.eat()
	mt.eat()
}


// The general form for type assertion expression is given as follows:
// <interface_variable>.(concrete type name)
func eat(f food) {
	// value, boolean := <interface_variable>.(concrete type name)
	// a much nicer idiom for assertions is the type switch statement. 
	// It uses the switch statement semantic to query static type 
	// information from an interface value using case clauses.
	switch morsel := f.(type) {
	case veggie:
		if morsel == "okra" {
			fmt.Println("Yuk! not eating ", morsel)
		} else {
			morsel.eat()
		}
	case meat:
		if morsel == "beef" {
			fmt.Println("Yuk! not eating ", morsel)
		} else {
			morsel.eat()
		}
	default:
		fmt.Println("Not eating whatever that is: ", f)
	}
}



func printAnyType(value interface{}) {
	fmt.Println(value)
}
