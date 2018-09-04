//
// Composite Types
//
package main

import(
	"fmt"
	"time"
	"math/rand"
	"encoding/json"
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
	historgram1 []map[string]int // slice of maps
	graph = [][][][]int{
		{{{44}, {3, 51}}, {{55, 12, 3}, {22, 4}}},
		{{{22, 12, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},
	}
)

//
// In general, map type is specified as follows:
// map[<key_type>]<element_type>
//
// The key specifies the type of a value that will be used 
// to index the stored elements of the map. 
// Map keys, MUST be of types that are comparable 
// including numeric, string, Boolean, pointers, arrays, struct, 
// and interface types (see Chapter 4 , Data Types, for 
// discussion on comparable types).
var (
	legends map[int]string
	histogram_map map[string]int = map[string]int{
		"Jan":100, "Feb":445, "Mar":514, "Apr":233,
		"May":321, "Jun":644, "Jul":113, "Aug":734,
		"Sep":553, "Oct":344, "Nov":831, "Dec":312,
	}
	calibration map[float64]bool
	matrix_map map[[2][2]int]bool // map whith array key type
	table map[string][]int = map[string][]int{ // map of string:[]integer slices as K:V
		"Men": []int{32, 55, 12, 58, 42, 76},
		"Women": []int{44, 23, 43, 65, 38, 51},
	}
	log map[struct{name string}]map[string]string
)



//
// The struct type
// The last type discussed in this chapter is Go's struct. 
// It is a composite type that serves as a container 
// for other named types known as fields.
// Note that the struct type has the following general format:
// struct{<field declaration set>}
// Accessing struct fields
// A struct uses a selector expression (or dot notation) 
// to access the values stored in fields.
// 	person.name
// 	person.address.street
// 	person.address.city
var (
	empty struct{}
	car = struct{mark, model string}{mark: "Ford", model: "F150"}
	currency = struct{name, country string; code int}{
		"USD", "United States",
		840,
	}
	node = struct{
		edges []string
		weight int
	}{
		edges: []string{"north", "south", "west"},
	}
	person struct{
		name string
		address struct{
			street string
			city, state string
			postal string
		}
	}
)

// Declaring named struct types
// Attempting to reuse struct types can get unwieldy fast.
type person_t struct {
	name string
	address address_t
}

type address_t struct {
	street string
	city string
	state string
	postal string
}

// The anonymous fields
// -The name of the type becomes the name of the field
// -The name of an anonymous field may not clash with other field names
// -Use only the unqualified (omit package) type name of imported types
type diameter int
type name struct {
	long string
	short string
	symbol rune
}
type planet struct {
	diameter
	name
	desc string
}



type person2_t struct {
	name string
	title string
}

// Field tags
// The following shows a definition of the Person and Address i
// structs that are tagged with JSON annotation which can be 
// interpreted by Go's JSON encoder and decoder (found in the
// standard library)
type Person struct {
	Name string `json:"person_name"`
	Title string `json:"perosn_title"`
	Address `json:"person_address_obj"`
}
// Notice the tags are represented as raw string values 
// (wrapped within a pair of `` ). The tags are ignored by 
// normal code execution. However, they can be collected 
// using Go's reflection API as is done by the JSON library.
type Address struct {
	Street string `json:"person_addr_street"`
	City string `json:"person_city"`
	State string `json:"person_state"`
	Postal string `json:"person_postal_code"`
}




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
	fmt.Println(" len: ", len(mnt), "cap: ", cap(mnt), mnt)
	mnt = append(mnt, "Jun", "Jul", "Aug")
	fmt.Println(" len: ", len(mnt), "cap: ", cap(mnt), mnt)

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

	hist := make(map[string]int)
	hist["Jan"] = 100
	hist["Feb"] = 445
	hist["Mar"] = 514
	fmt.Println(" hist: ", hist)
	jan_val := hist["Jan"]
	fmt.Println(" jv: ", jan_val)
// uncomment for testing!
	//save(hist, "Feb", 1334)

	// map traversal
	// Each iteration returns a key and its associated element value. 
	// Iteration order, however, is not guaranteed.
	fmt.Println("[Histogram]")
	for k,v := range histogram_map {
		adjVal := int(float64(v) * 0.1000)
		fmt.Printf(" %s (%d): ", k, v)
		for i := 0; i < adjVal; i++ {
			fmt.Print(".")
		}
		fmt.Println()
	}

	// map functions
	fmt.Println(" before len: ", len(hist)) // map contains 3 item
	var err error = remove(hist, "Jun")
	if err != nil {
		// here we have got an error, because there is no key "Jun"
		fmt.Println("E: ", err)
	}
	err = remove(hist, "Jan")
	if err != nil {
		// but here we are ok, there is "Jan" key which has been successfully found
		// no error 
		fmt.Println("E: ", err)
	}
	fmt.Println(" after len: ", len(hist)) // already, map contains 2 intems

	//
	// structs
	//
	fmt.Println(" car: ", car)
	fmt.Println(" cur: ", currency)
	fmt.Println(" node: ", node)
	fmt.Println(" person: ", makePerson())
	earth := planet{
		diameter: 7926,
		name: name{
			long: "Earth",
			short: "E",
			symbol: '\u2641',
		},
		desc: "Third rock from the Sun",
	}
	fmt.Println(" earth: ", earth)
	jupiter := planet{}
	jupiter.diameter = 88846
	jupiter.name.long = "Jupiter"
	jupiter.name.short = "J"
	jupiter.name.symbol = '\u2643'
	jupiter.desc = "A ball of gas"
	fmt.Println(" jupiter: ", jupiter)
	// Promoted fields
	// this will only work if the promotion does not cause any identifier clashes. 
	// In case of ambiguity, the fully qualified selector expression can be used.
	saturn := planet{}
	saturn.diameter = 120536
	saturn.long = "Saturn" // promoted field
	saturn.short = "S" // promoted field
	saturn.symbol = '\u2644' // promoted field
	saturn.desc = "Slow mover"
	fmt.Println(" saturn: ", saturn)
	// sturct as params
	p := new(person2_t)
	p.name = "undefined"
	// or we can use &(ampersand) to get an address if p is an object of the struct persone2_t
	fmt.Println(" before p: ", p)
	updateName(p, "Barmosnatch")
	fmt.Println(" after p: ", p)

	p1 := Person{}
	p1.Name = "Faldy Follen"
	p1.Title = "Author"
	p1.Street = "415th st"
	p1.City = "Govile"
	p1.State = "Gonecticut"
	p1.Postal = "453456"
	encP1,_ := json.Marshal(p1)
	fmt.Println(" jsoned: ", string(encP1))
}

func updateName(p *person2_t, name string) {
	p.name = name
}



func makePerson() person_t {
	addr := address_t {
		city: "Goville",
		state: "Gonecticut",
		postal: "342312",
	}
	return person_t {
		name: "Circul Schturman",
		address: addr,
	}
}

func remove(store map[string]int, key string) error {
	_,ok := store[key]
	if !ok {
		return fmt.Errorf("Key not foud")
	}
	delete(store, key)
	return nil // no error occured
}

func save(store map[string]int, key string, value int) {
	// Called the comma-ok idiom, the Boolean value stored in 
	// the ok variable is set to false when the value is not actually found.
	v, ok := store[key]
	if !ok {
		store[key] = value
	} else {
		panic(fmt.Sprintf("Slot %d taken ", v))
	}
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

