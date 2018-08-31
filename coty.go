//
// Composite Types
//
package main

import(
	"fmt"
	"time"
	"math/rand"
)

var val [100]int = [100]int{44,72,12,55,64,1,4,90,13,54}
var days [7]string = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

var truth = [256]bool{true}
var histogram = [5]map[string]int{
	map[string]int{"A":12,"B":1,"D":15},
	map[string]int{"man":1344, "women":844, "children":577},
}

var board = [4][2]int{
	{33,23},
	{62,2},
	{23,4},
	{51,88},
}

var matrix = [2][2][2][2]byte{
	{{{4,4}, {3,5}}, {{55,12}, {22,4}}},
	{{{2,2}, {7,9}}, {{43,0,}, {88,7}}},
}

var weekdays = [...]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thuesday",
	"Friday",
}

var msg = [12]rune{0:'H', 2:'E', 4:'L', 6:'O', 8:'!'}

type matrix_t [2][2][2][2]byte

const size = 1000
var nums [size]int

type numbers_t [1024 * 1024]int

type galaxies_t [14]string



//
// main driver
//
func main() {
	fmt.Println("composit")
	mx := initMx()
	fmt.Println("matrix: ", mx)
	var seven = [7]string{
		"grumpy",
		"sleepy",
		"bashful",
	}
	// determinig length and capacity functions
	fmt.Println(len(seven), cap(seven))

	initRandArray()
	fmt.Println("max: ", maxRandArray(nums))

	var a numbers_t
	initRandArrayParamPtr(&a)
	fmt.Println("max2: ", maxRandArrayParamPtr(&a))

	namedGalaxies := &galaxies_t{
		"Andromeda",
		"Black Eye",
		"Body's",
		"Cassiopeia",
	}
	printGalaxies(namedGalaxies)
}

func printGalaxies(galaxies *galaxies_t) {
	fmt.Println(" [Galaxies]")
	for _, g := range galaxies {
		fmt.Println(" :", g)
	}
}

func initMx() matrix_t {
	return matrix_t{
		{{{4,4}, {3,5}}, {{55,12}, {22,4}}},
		{{{2,2}, {7,9}}, {{43,0,}, {88,7}}},
	}
}


func initRandArray() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(10000)
	}
}

func maxRandArray(nums [size]int) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

//
// -------------------------------
//

func initRandArrayParamPtr(nums *numbers_t) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(10000)
	}
}

func maxRandArrayParamPtr(nums *numbers_t) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}

