// load main method package
package main

import (
	"fmt"
	"math/cmplx"
)

// simple addition function
// note that type declaration is after var names
// and if multiple vars have same type, only declare it once at the end
func add(x, y int) int {
	return x + y
}

// simple swap string function
// note you can return multiple output vars
func swap(x, y string) (string, string) {
	return y, x
}

// naked return swap function
// note that you simply specify return without vars
// it understands that you are returning x and y due to function definition at top
// only use this in short functions or it harms readability
func nakedCalc(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// var allows creation of 1 or more vars with type
// type can be omitted and inferred from variable assignment
// var can be defined outside functions (here) or inside (below)
var i, j int = 1, 2

// simple function to show named vars
func showVars(i, j int) (int, int, string, bool) {
	// var types are inferred from short variable declaration
	// you can't use := outside a function
	str1, str2 := "test", true

	// return input values and local vars
	return i, j, str1, str2
}

// variable assignment can be put into blocks
// many different var types in Go (e.g. complex)
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// run main method
func main() {
	// call add function
	fmt.Println(add(3, 8))

	// swap strings (simple var assignment)
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// call nakedCalc
	fmt.Println(nakedCalc(17))

	// call showVars
	fmt.Println(showVars(4, 5))

	// show var types
	fmt.Println(ToBe, MaxInt, z)

	// vars default to specific values if not defined
	// int & float > 0, bool > false, str > ""
	// constants are similar to variables but their value never changes
	// numeric constants are high precision numbers (i.e. 1x10-20)
	var i int
	var f float32
	var bo bool
	var s string
	const Hello = "hello"
	fmt.Printf("%v %v %v %q %q\n", i, f, bo, s, Hello)

	// var types can be reassigned as follows
	// note that you cannot add vars of different type (even uint and int)
	// reason for this here: https://blog.golang.org/constants
	// essentially don't want bugs due to type mismatches
	var ij int = 5
	var jk float64 = float64(ij)
	var kl uint = uint(jk) // uint = unsigned = cannot be negative
	fmt.Println(ij, jk, kl)
}