package main

import (
	"fmt"
	"math"
)

// simple function to return root of value
func sqrt(x float64) string {
	// simple if statement
	// if value < 0
	if x < 0 {
		// re-run function on positive value
		return sqrt(-x) + "i"
	}

	// if positive, simply get root of value
	return fmt.Sprint(math.Sqrt(x))
}

// function to check if product of 2 int vars exceeds limit
func limCalc(x, y, lim int) bool {
	// you can use short assignment within an if statement declaration
	// multiple if statements can be chained using ;
	// note that v cannot be used outside of the if statement, it is local to that chunk
	if v := x * y; v < lim {
		return true
	} else {
		// we can still reference v here as an else block
		// is technically still part of the if statement
		fmt.Printf("Can use v here: %v\n", v)
		return false
	}
	// cannot reference v from here onwards
	// now that we're outside of the if statement
}

// switch statements can check specific conditions
// below we have a condition of 10
// meaning that every case checks to see if it equals 10
// this can be used for things like checking dates, times, results of functions etc.
func checkVal2(x int) {
	switch 10 {
	case x:
		fmt.Println("x = 10")
	default:
		fmt.Println("x does not equal 10")
	}
}

// switch statements make if/else simpler
func checkVal(x, y, z int64) {
	// run each case and return outputs conditionally
	// switch cases always run from top to bottom
	// default will run if no cases are met
	// switch statements without conditions are simply if/else chains
	// i.e. below there is no variable directly after the switch statement
	switch {
	case x > y:
		fmt.Printf("%v > %v\n", x, y)
	case x > z:
		fmt.Printf("%v > %v\n", x, z)
	default:
		fmt.Printf("%v > %v and %v \n", x, y, z)
	}
}

// function to find square root arithmetically
// using Newton's law (derivatives)
// x, y are starting guess of sqrt and desired value (squared value)
// z is number of iterations to try
func findSqrt(x, y float64, z int) float64 {
	// track x before/after increment
	var xNew float64 = x

	// repeat for z iterations
	for i := 1; i <= z; i++ {
		// store current version of x
		var xOld = xNew

		// calculate distance between square of x and actual y
		// scale down difference using derivative function (i.e. / 2x)
		// decrement x to get closer to actual square root
		xNew -= ((xNew * xNew) - y) / (2 * xNew)

		// show progress
		fmt.Println(xNew, xNew * xNew, y)

		// return output if value has converged (i.e. found perfect square root)
		if xNew - xOld == 0 {
			// return final approximation of sqrt
			return xNew
		}
	}

	// return final approximation of sqrt
	return x
}

// defer allows you to store values/actions
// and then run them once the rest of the function has completed
// more info: https://blog.golang.org/defer-panic-and-recover
func deferFlow() {
	// this value is stored, but not executed
	// until rest of function has completed
	defer fmt.Println("World")

	// defer operates in a last in first out (LIFO) method
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	// this line will be run first
	fmt.Println("Hello")
}

func main() {
	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// while loop (just for without semi-colons)
	i := 1
	for i < 10 {
		i += i
	}

	// print output
	//fmt.Println(sum)

	// call sqrt function
	//fmt.Println(sqrt(16), sqrt(-4))

	// call limCalc function
	//fmt.Println(limCalc(5, 4, 18))

	// call sqrt calculator
	// define input vars
	var (
		x, y, z = 1.0, 16.0, 10
	)
	// calculate approximated square root
	approx := findSqrt(x, y, z)
	// show results
	fmt.Printf("Starting with %v, the sqrt of %v is %v after %v iterations.\n", x, y, approx, z)

	// run checkVal
	checkVal(2, 4, 6)
	checkVal(8, 9, 6)

	// run checkVal2
	checkVal2(10)
	checkVal2(5)

	// run deferFlow
	deferFlow()
}
