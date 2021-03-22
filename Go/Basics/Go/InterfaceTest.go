// load packages
package main

// import libs
import (
	"fmt"
	"math"
)

// create interface for Shape object
// implements Area and Perimeter methods (defined below)
// interfaces can handle nil values (i.e. no input when called) without raising an error
type Shape interface {
	Area() float64
	Perimeter() float64
}

// create rectangle struct
// contains width and height fields
type Rect struct {
	width float64
	height float64
}

// create circle struct
// contains radius field
type Circle struct {
	radius float64
}

// define area method for rectangles
// note how the name of this method is identical to the below equivalent for circles
// this is fine, as the Shape interface will dynamically interpret the loaded object
// this dynamic handling is known as polymorphism and it happens implicitly in Go
func (r Rect) Area() float64 {
	return r.height * r.width
}

// define perimeter method for rectangles
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// define area method for circles
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius // pi * r squared
}

// define perimeter method for circles
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius // 2 * pi * r
}

// main method
func main() {
	// create and assess rectangle object via shape interface
	// note that you can access the methods within the shape interface (e.g. Area(), Perimeter())
	// but you cannot access the fields within the shapes themselves (e.g. radius, width, height)
	// note also that you can unpack the value and the type of an interface object, the type is a concrete type
	var s Shape = Rect{3, 6}
	fmt.Printf("This shape is a %T\n", s)
	fmt.Printf("This perimeter of this %T is %v\n", s, s.Perimeter())
	fmt.Printf("This area of this %T is %v\n", s, s.Area())

	// create and access circle object via shape interface
	// note how the Shape interface dynamically interprets the methods to use
	// based upon the struct type loaded into the interface
	// here, we can calculate perimeter using the circle specific method because
	// the object created is a circle and not a rectangle
	s = Circle{5}
	fmt.Printf("This shape is a %T\n", s)
	fmt.Printf("The perimeter of this %T is %v\n", s, s.Perimeter())
	fmt.Printf("The area of this %T is %v\n", s, s.Area())
}