//
// The break, continue, and goto statements
//
package main

import(
	"fmt"
)

var words = [][]string{
	{"break", "lake", "go", "right", "strong", "kite", "hello"},
	{"fix", "river", "stop", "left", "weak", "flight", "bye"},
	{"fix", "lake", "slow", "middle", "sturdy", "high", "hello"},
}


func searchOnce(w string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == w {
				fmt.Println("Found w -> ", w)
				break DoSearch
			}
		}
	}
}


func searchAll(w string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == w {
				fmt.Println("Found w -> ", w)
				continue DoSearch
			}
		}
	}
}

func main() {
	searchOnce("slow")
	searchAll("lake")

	// goto usage example FSM-like style
	var a string
	Start:
		for {
			switch {
			case a < "aaa":
				fmt.Println(" CONTINUE -> A")
				goto A
			case a >= "aaa" && a < "aaabbb":
				fmt.Println(" CONTINUE -> B")
				goto B
			case a == "aaabbb":
				fmt.Println(" BREAK -> Start")
				break Start
			} //eof switch
		A:
			a += "a"
			continue Start
		B:
			a += "b"
			continue Start
		} //eof for
	fmt.Println(" Result -> ", a)
}
