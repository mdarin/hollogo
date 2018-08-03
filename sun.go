//
// Go variable declaration and initialization example
//
package main

import "fmt"
import "time"

// declare variables that describes star info
var name string
var desc string
var radius int32
var mass float64
var active bool
var satellites []string

func main() {

	// init start info
	name = "Sun"
	desc = "Star"
	radius = 685800
	mass = 1.989E+30
	active = true

	// fillin the sattellits' array 
	satellites = []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// show star info	
	fmt.Println("Name ", name)
	fmt.Println("Desc ", desc)
	fmt.Println("Active ", active)
	fmt.Println("Radius (km)", radius)
	fmt.Println("Mass (kg)", mass)
	fmt.Println("Statellites ", satellites)

	fmt.Println("-----------------------------------")

	// Restrictions for short variable declaration
	// For convenience, the short form of the variable declaration does come with several
	// restrictions that you should be aware of to avoid confusion:
	// Firstly, it can only be used within a function block
	// The assignment operator := , declares variable and assign values
	// := cannot be used to update a previously declared variable
	// Updates to variables must be done with an equal sign

	// using the short variable declaration
	name_n := "Neptune"
	desc_n := "Planet"
	radius_n := 24764
	mass_n := 1.024e26
	active_n := true
	satellites_n := []string{
		"Naiad", "Thalassa", "Despina", "Galatea", "Larissa",
		"S/2004 N 1", "Proteus", "Triton", "Nereid", "Halimede",
		"Sao", "Laomedeia", "Neso", "Psamathe", //the last one comma required!
	}

	// show info	
	fmt.Println("Name ", name_n)
	fmt.Println("Desc ", desc_n)
	fmt.Println("Active ", active_n)
	fmt.Println("Radius (km) ", radius_n)
	fmt.Println("Mass (kg) ", mass_n)
	fmt.Println("satellites ", satellites_n)

	fmt.Println("-----------------------------------")

	var (
		name_e string = "Eath"
		desc_e string = "Planet"
		radius_e string = "6378"
		mass_e float64 = 5.974E+24
		active_e bool = true
		satellites_e []string
	)

	// show info	
	fmt.Println("Name ", name_e)
	fmt.Println("Desc ", desc_e)
	fmt.Println("Active ", active_e)
	fmt.Println("Radius (km) ", radius_e)
	fmt.Println("Mass (kg) ", mass_e)
	fmt.Println("satellites ", satellites_e)

	fmt.Println("--------------constatnts-------------------")
	const c rune = 'G'
	const a string = "Mastering"
	const b string = "Go"
	//
	// One advantage of using untyped constants is that the 
	// type system relaxes the strict application of type checking.
	//
	const o = time.Millisecond * 5
	//тут ошибка возинкает...const k1, k2 = true, !k1
	const(
		h = time.Second * 4
	)

	fmt.Println("Go constant value")
	fmt.Println("typed const a", a)
	fmt.Println("typed const b", b)
	fmt.Println("typed const c", c)
	//fmt.Println("untyped const k1", k1)
	//fmt.Println("untyped const k2", k2)
	fmt.Println("untyped const o", o)
	fmt.Println("const block h", h)

	//
	// Constant enumeration
	//
	// The compiler will then automatically do the following:
	// -Declare each member in the block as an untyped integer constant value
	// -Initialize iota with a value of zero
	// -Assign iota , or zero, to the first constant member ( StarHyperGiant )
	// -Each subsequent constant is assigned an int value increased by one
	const (
		StarHyperGiant = iota
		StarSuperGiant
		StarBrightGiant
		StarGiant
		StarSubGiant
		StarDwarf
		StarSubDwarf
		StarWhiteDwarf
		StarRedDwarf
		StarBrownDwarf
	)
	

}

