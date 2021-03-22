package main

import (
	"errors"
	"fmt"
	"math"
)

// Go doesn't have classes
// instead it uses functions with input vars, called methods
// first, let's create a struct which will be an input for a method
type Vertex struct {
	X, Y float64
}

// let's create a method called Abs()
// which takes a Vertex struct as an input and returns an int
// note that methods can take inputs of any type (float, int etc.)
// but the input type must be defined within the same package, it cannot be external
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// a method is just a function with a receiver element in the declaration line
// the below is identical to the above, except it's a pure function instead of a method
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// methods can also take pointers as inputs (a.k.a. pointer receivers)
// here, because we use a pointer to v, we are modifying the existing Vertex
// if we omitted the pointer/* and just used a new v Vertex input
// we would be performing the below steps to an entirely separate copy of v
// methods are clever, and understand that if the receiver is a pointer, it must
// invoke & to create a point (i.e. it infers v.X as (&v).X to create a pointer)
// plain functions do not do this, they will raise errors if you don't explicitly use &
// NOTE: it is best to use pointers where possible as it doesn't create multiple copies of a var (i.e. memory efficient)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// interfaces can be created without specified types/methods inside them
// in this case, they can take values of any type
// fmt.Print is an example of this, it will happily print multiple types of input
// function to show value and type of interface item
func describe(i interface{}) {
	fmt.Printf("Type: %T and Value: %v\n", i, i)
}

// function to run multiple type assertions using a switch statement
func checkType(i interface{}) {
	// create instance of interface and check its type (note: type is a keyword here)
	switch v := i.(type) {
	// check if it's an int
	case int:
		fmt.Printf("Type: %T and Value: %v\n", v, v)
	// check if it's a float
	case float64:
		fmt.Printf("Type: %T and Value: %v\n", v, v)
	// if neither of the above, default to below and return whatever it actually is
	default:
		fmt.Printf("Type: %T and Value: %v\n", v, v)
	}
}

// stringers are functions/methods which format and return string variables
// first, let's create a variable for a stringer to process
type IPAddr [4]byte

// now let's create a function which takes in the above variable
// and formats it to show each value separated by a full stop
func (ip IPAddr) String() string {
	// Sprintf formats the elements of our input variable by putting a full stop between each value
	// it then returns this value (this is what Sprintf does,
	// whereas Printf for example returns nothing, it simply prints)
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

// ERROR HANDLING
// let's create a simple divide function that takes 2 ints as an input
// and returns an int if they can be divided or an error if not
func Divide(a int, b int) (int, error) {
	// cannot divide by 0
	if b == 0 {
		// use in-built errors interface to return error message
		return 0, errors.New("cannot divide by 0")
	}
	// otherwise perform calculation
	return (a / b), nil
}

func main() {
	// METHODS AND POINTERS
	// create an instance of a Vertex
	v := Vertex{3, 4}

	// apply scaling method (this will modify v as it's a pointer)
	v.Scale(10)

	// apply method to Vertex
	fmt.Println(v.Abs())

	// EMPTY INTERFACES
	// prove empty interfaces can handle multiple types
	// these are useful for when you don't know what type your returned variables will be (e.g. fmt.Print)
	var i interface{} // create interface with no specified types
	describe(i) // show nil contents
	i = 1
	describe(i) // interface handles int type
	i = "Hello"
	describe(i) // it also handles strings

	// TYPE ASSERTIONS
	// create an interface containing a string
	var j interface{} = "Hello"
	t, ok := j.(string) // check if interface value j contains a string (t will be the type and ok will be true/false)
	fmt.Println(t, ok)
	//t, ok = j.(float64) // check if interface value j contains a float (it doesn't, so it will panic)
	//fmt.Println(t, ok)

	// SWITCH & TYPE ASSERTIONS
	checkType(21)
	checkType("Hello")
	checkType(true)

	// STRINGERS
	// now we can provide inputs of type IPAddr (i.e. 4 byte values)
	// and our stringer method (i.e. string format and return method) will handle it for us
	ip := IPAddr{127, 0, 0, 1}
	fmt.Printf(ip.String())

	// ERROR HANDLING
	// errors should be nil if everything has worked
	// if they're not nil, the error should be called and shown
	fmt.Println(Divide(5, 2))
	fmt.Println(Divide(8, 0))
	if result, error := Divide(5, 2); error != nil {
		fmt.Println("Error occurred: ", error)
	} else {
		fmt.Println(result)
	}
}
