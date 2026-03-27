package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

func (r Circle) Perimeter() float64 {
	return 2 * math.Pi * r.Radius
}

func interfaces() {
	// Interfaces in Go define a set of method signatures. Any type that implements all these
	// methods automatically satisfies the interface. no explicit declaration required.
	// This is called implicit satisfaction.

	// Interfaces are usually named with words ending with "-er".

	// From the types and their respective methods, we can see that both the Rectangle and Circle
	// types, satisfy the Shaper interface.

	printShapeInfo(Rectangle{
		Width:  12.54,
		Height: 9.38,
	})
	printShapeInfo(Circle{
		Radius: 9.99,
	})

	var s Shaper = Rectangle{Width: 12, Height: 3}

	fmt.Printf("%T\n", s)

	//* The empty interface
	// `any` is an alias for `interface{}`, and it has no methods. This interface will accept every type.

	printAnything("wtf")
	printAnything(920)
	printAnything(Circle{Radius: 31})

	//* Embedding and Interfaces
	// Intrerfaces can also be embedded into other interfaces,

	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	type Closer interface {
		Close() error
	}

	type ReadCloser interface {
		Reader
		Closer
	}

	//* Accept Interfaces, return Structs
  
}

func printShapeInfo(s Shaper) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func printAnything(v any) {
	fmt.Println(v)
}
