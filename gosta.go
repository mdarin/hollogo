// 
// go control flow
//
package main

import (
	"fmt"
	"strings"
)

type Curr struct {
	Currency string
	Name string
	Country string
	Number int
}

var currencies = []Curr{
	Curr{"DZD", "Algerian Dinar", "Algeria", 12},
	Curr{"AUD", "Australian Dollar", "Australia", 36},
	Curr{"EUR", "Euro", "Belgium", 978},
	Curr{"CLP", "Chilean Peso", "Chile", 152},
//...
}


//
// main driver
//

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

	// how to use more complex function 
	find("EUR")

	if	assertEuro(Curr{"EUR", "Euro", "Belgium", 978}) {
		fmt.Println("is a euro currency")
	} else {
		fmt.Println("is another currency")
	}

	findNumber(36)

	findAny(404)
	findAny(978)
	findAny("AUD")
	findAny(false)

	// show unsorted
	listCurr()
	// sort
	sortByNumber()
	// show sorted
	listCurr()

}

func listCurr() {
	// like a while(true)
	i := 0
	for i < len(currencies) {
		//FIXME: how does it work correct, humm?
		fmt.Printf(" %d: %s\n",i,currencies[i])
		i++
	}
}


//my firts function :) 
func myFunc() bool {
	return true
}


func find(name string) {
	for i := 0; i < len(currencies); i++ {
		c := currencies[i]
		switch {
		case strings.Contains(c.Currency, name),
			strings.Contains(c.Name, name),
			strings.Contains(c.Country, name): fmt.Println("found -> ", c)
		}
	}
}


func assertEuro(c Curr) bool {
	// Switch initializer example
	// Notice the trailing semi-colon to indicate the separation between 
	// the initialization statement and the expression area for the switch.
	// In the example, however, the switch expression is empty.
	switch name, curr := "Euro", "EUR"; {
		case c.Name == name:
			return true
		case c.Currency == curr:
			return true
	}
	return false
}


// Go offers the type interface{} , or empty
// interface, as a super type that is implemented by all other types in the type system. When a
// value is assigned type interface{} , it can be queried using the type switch , as shown in
// function findAny() in the following code snippet, to query information about its
// underlying type

func findNumber(num int) {
	for _, curr := range currencies {
		if curr.Number == num {
			fmt.Println("found2 -> ", curr)
		}
	}
}


func findAny(val interface{}) {
	switch i := val.(type) {
	case int: findNumber(i)
	case string: find(i)
	default: fmt.Printf("Unable to serach with type %T\n", val)
	}
}



func sortByNumber() {
	N := len(currencies)
	for i := 0; i < N-1; i++ {
		currMin := i
		for k := i + 1; k < N; k++ {
			if currencies[k].Number < currencies[currMin].Number {
				currMin = k
			}
		}
		// swap
		if currMin != i {
			temp := currencies[i]
			currencies[i] = currencies[currMin]
			currencies[currMin] = temp
		}
	}
}


