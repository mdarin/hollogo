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

// the slice type and initialization
var(
	image []byte
	ids []string = []string{"fe225", "ac144", "3d12c"}
	vector = []float64{12.4, 44, 126, 2, 11.5}
	monthsAsPtr = &[]string{
		"January",
		"Febuary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"Augest",
		"September",
		"October",
		"November",
		"December",
	}
	months = []string{
		"January",
		"Febuary",
		"March",
		"April",
		"May",
		"June",
		"July",
		"Augest",
		"September",
		"October",
		"November",
		"December",
	}
	tables = []map[string][]int {
		{
			"age": {53, 13, 5, 55, 45, 62, 34, 7},
			"pay": {124, 66, 777, 531, 933, 231},
		},
	}
	q1 []string
	historgram []map[string]int // slice of maps
	graph = [][][][]int{
		{{{44}, {3, 51}}, {{55, 12, 3}, {22, 4}}},
		{{{22, 12, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},
	}
)

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
	// gethering address of months array
	printMonths(&months)
	// using ptr
	printMonths(monthsAsPtr)

	// slicing
	allM := months[:]
	m1 := months[:3]
	m2 := months[3:]
	mappedM3 := months[2:4]
	fmt.Println(" all m: ", allM)
	fmt.Println(" m1: ", m1)
	fmt.Println(" m2: ", m2)
	fmt.Println(" m3: ", mappedM3)
	// expressions with capacity
	// <slice_or_array_value>[<low_index>:<high_index>:max]
	su := months[5:8:8]
	fmt.Println(" summer: ", su)

	// make() function
	// making a slice
	// This snippet will initialize the variable with a slice value with an initial
	// length of 6 and a maximum capacity of 12
	slice1 := make([]string, 6, 12)
	fmt.Println(" slice: ", slice1)

	vector = scale(2.45, vector)
	fmt.Println(" new v: ", vector)
	var cont bool = contains(float64(4.9), vector)
	fmt.Println(" cont: ", cont)
	// slices length and capacity
	var vt []float64
	fmt.Println(" zero-len: ", len(vt))
	h := make([]float64, 4, 10)
	fmt.Println(" len: ", len(h), ", cap: ", cap(h))
	mnt := make([]string, 3, 3)
	mnt = append(mnt, "Jan", "Feb", "March")
	fmt.Println("len: ", len(mnt), "cap: ", cap(mnt), mnt)
	mnt = append(mnt, "Jun", "Jul", "Aug")
	fmt.Println("len: ", len(mnt), "cap: ", cap(mnt), mnt)

	// copying
	cp := clone(vector)
	fmt.Println(" copy: ", cp)

	// stirng as skuces
	msg := "Bobsayshelloworld!"
	fmt.Println(
		" splitted: ",
		msg[:3], msg[3:7], msg[7:12],
		msg[12:17], msg[len(msg)-1:],
	)
	var sorted string = sort(msg)
	fmt.Println(" sorted: ", sorted)
}


func sort(str string) string {
	// The code shows the explicit conversion of a slice of bytes to a string value. Note
	// that each character may be accessed using the index expression.
	bytes := []byte(str)
	var temp byte
	// sorting...
	for i := range bytes {
		for j := i + 1; j < len(bytes); j++ {
			if bytes[j] < bytes[i] {
				temp = bytes[i]
				bytes[i], bytes[j] = bytes[j], temp
			}
		}
	}
	return string(bytes)
}


// The copy function copies the content of v slice into result. 
// Both source and target slices must be the same size and of the
// same type or the copy operation will fail.
func clone(v []float64) (result []float64) {
	result = make([]float64, len(v), cap(v))
	copy(result, v)
	return
}

func scale(factor float64, vector []float64) []float64 {
	for i := range vector {
		vector[i] *= factor
	}
	return vector
}

func contains(val float64, vector []float64) bool {
	for _, num := range vector {
		if num == val {
			return true
		}
	}
	return false
}


func printMonths(months *[]string) {
	fmt.Println("[Months]")
	for _, m := range *months {
		fmt.Println(" m: ", m)
	}
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

