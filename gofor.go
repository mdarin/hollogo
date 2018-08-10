//
// The for range
//
package main

import(
	"fmt"
	"math/rand"
)


var list1 = []string{
	"brea","lake","go",
	"right","strong",
	"kite","hello",
}

var list2 = []string{
	"fix","river","stop",
	"left","weak","flight",
	"bye",
}


//
// Main driver
//

func main() {
	rand.Seed(31)
	for w1, w2 := nextPair(); w1 != "go" && w2 != "stop"; w1,w2 = nextPair() {
		fmt.Printf("Wrod pair -> [%s, %s]\n", w1, w2)
	}

	vals := []int{33,44,55}

	// just show vals, 
	for _,v := range vals {
		fmt.Println(v)
		v--
	}
	fmt.Println(vals)

	// modify vals
	for i,v := range vals {
		vals[i] = v - 1
	}
	fmt.Println(vals)

	// The next form of the for statement was introduced (as of Version 1.4 of Go) 
	// to express a for range without any variable declaration
	for range []int{1,2,1,2,3,4} {
		fmt.Println("Looping")
	}
}


func nextPair() (w1, w2 string) {
	pos := rand.Intn(len(list1))
	return list1[pos],list2[pos]
}

