package main

import (
	"fmt"
	"math"
)

// structs are ways of keeping multiple fields in one place
// they're like a mini/simple class in some ways
type person struct {
	name string
	age int
}

// struct literals are values assigned to struct fields
// upon creation, they can be omitted for default values
type vertex struct {
	x, y int
}

// MAPS
// maps are like python dictionaries, allowing key and value storage
// first let's create a struct which will be our values
type Vertexx struct {
	Lat, Long float64
}

// now let's define our map
//var m map[string]Vertex // the format here is map[key type]value type

// we use make to instantiate an empty map
var m = make(map[string]Vertexx)

// as with most objects, you can create map literals (i.e. instantiate with values)
var ma = map[string]Vertex{
	"First Key": Vertex{25.32, 39.04},
	"Second Key": Vertex{-75.24, -11.52},
}

func maps() {
	// and populate with struct values
	m["First Key"] = Vertexx{33.46, 42.18}

	// access map elements
	fmt.Println(m["First Key"])

	// using/mutating maps
	m["First Key"] = Vertexx{1.1, 2.2} // insert/update value at element
	var val = m["First Key"] // assign map value to var
	delete(m, "First Key") // delete specific key from map
	var ok bool
	val, ok = m["First Key"] // extracts value to val if present (otherwise 0) and sets ok to true if present, false if not

	// use vars
	fmt.Println(val, ok)
}

// FUNCTIONS
// functions can be used as arguments to other functions
// and can also be returned as outputs of other functions
// this function returns a float after taking an input of another function which takes 2 floats in and returns a float
func compute(fn func(float64, float64) float64) float64 {
	// evaluate input function using the values 3 and 4
	return fn(3, 4)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	// function within a function
	fmt.Println(hypot(5, 7)) // simply runs hypot
	fmt.Println(compute(hypot)) // runs hypot using values within compute
	fmt.Println(compute(math.Pow)) // runs x**y using values within compute

	// 1) pointers
	// create standard int vars
	var i, j int = 10, 15

	// create pointers
	// this is known as de-referencing or indirecting
	// & generates a point, * references a pointer's value
	p := &i // create pointer p that points to i
	fmt.Println(*p) // read i via pointer
	*p = 21 // set i through pointer
	fmt.Println(i) // confirm that value of i has changed

	p = &j // point to new var
	*p = *p * 4 // carry out var arithmetic via p
	fmt.Println(j) // prove that pointer affected underlying var

	// 2) structs
	p1 := person{"Matt", 28} // create instance of struct
	point := &p1 // create pointer to struct instance
	fmt.Println(point.name) // access struct field
	point.age = 29 // update struct field name via pointer
	//(*point).age = 29 // this is identical to the above line but Go allows you to drop additional syntax
	fmt.Println(p1.age) // show that pointer allowed update of struct field

	// struct literals definition
	v1 := vertex{1, 2} // create struct with explicit values
	v2 := vertex{y:3} // specify y explicitly, x will default to 0
	v3 := vertex{} // x and y will default to 0
	v4 := &vertex{1, 2} // create a pointer to struct
	fmt.Println(v1, v2, v3, v4)

	// 3) arrays and slices
	// arrays length is set, you cannot change it
	// i.e. a (below) will always have 10 index positions
	var a [10]string // create an array of 10 integers
	a[0] = "Hello" // assign value to array index
	a[1] = "World"
	fmt.Println(a[0], a[1]) // access array components

	// primes is an array literal, because you've explicitly defined the values
	primes := [6]int{2, 3, 5, 7, 11, 13} // initialize with specified values

	// you can append values to a slice
	pr := primes[:]
	pr = append(pr, 17)

	// slices are subsections of arrays and are often used more than rigid arrays
	// unlike arrays, their length can be anything
	var s []int = primes[1:4] // create slice of primes array from index 1 to 3 (incl first, excl last)
	fmt.Println(s) // show slice components

	// slices are essentially views of arrays
	// they do not contain data themselves, just specific elements of the underlying array
	// as such, any changes you make to a slice will affect the underlying array and any other slices
	names := [4]string {"John", "Paul", "Ringo", "George"} // create fixed array of 4 strings
	var s1 []string = names[0:2] // extract first 2 names into slice
	s2 := names[1:3] // extract 2nd and 3rd names into slice (NOTE: identical to above, just shorthand code)
	//s2 := names[1:] // can omit values to begin at start or run to end
	//s2 := names[:3]
	s2[0] = "XXX" // change first item in s2 to "XXX" (NOTE: this changes underlying array names and s1 as well)
	fmt.Println(names, s1, s2) // show that change to s2 affects names and s1
	fmt.Println(len(s2), cap(s2)) // length is # of items in slice, capacity is # of items in underlying array

	// slice literals are like array literals except without a specified length
	// again, you instantiate them with specified values
	primeLit := []int{2, 3, 5, 7, 11, 13}
	structLit := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(primeLit)
	fmt.Println(structLit)

	// empty slices have nil value
	var ss []int
	fmt.Println(ss, len(ss), cap(ss)) // len and cap are 0
	if ss == nil {
		fmt.Println("nil!")
	}

	// make can be used to generate slices
	// this allows you to make dynamic slices
	// i.e. by passing variables where 0 and 5 are in the below line
	aa := make([]int, 5, 10) // create array with len 5 and cap 10
	bb := aa[:cap(aa)] // create array from index 0 to 10
	cc := aa[1:] // create array from index 1 to end (i.e. 5)
	fmt.Println(aa, bb, cc)

	// slices can contain any type, including other slices
	// create 2d slice containing other slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"0", "_", "X"},
		[]string{"_", "0", "X"},
	}

	// assign value to 2d position within slice of slices
	board[0][0] = "0"

	// show output
	fmt.Println(board)

	// range lets you iterate over a slice
	aaa := []int{1, 2, 3, 4, 5} // create array
	for i, v := range aaa {
		// print index and value within array
		fmt.Println(i, v)
	}

	// you can omit the index by providing an "_"
	for _, v := range aaa {
		fmt.Println(v)
	}

	// or you can omit the value by just requesting the index
	for i := range aaa {
		fmt.Println(i)
	}
}
