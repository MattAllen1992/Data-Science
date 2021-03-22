package main

import (
	"fmt"
	"time"
)

// simple function to print string 3 times with a count
func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, ": ", i)
	}
}

// main method
func main() {
	// call function normally
	// this will simply call the above function with the provided string
	f("First")

	// perform go routine
	// this will run concurrently with all other functions here (i.e. at the same time)
	go f("Second")

	// anonymous function within go routine
	// this will also run concurrently, interweaving with the above 2 functions
	go func(s string) {
		fmt.Println(s)
	} ("Third")

	// allow time for go routines to run
	// the output will show that the written order above is different to the
	// output order due to the multiple threads running concurrently and
	// simply returning values when complete, rather than waiting for previous steps to complete
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
