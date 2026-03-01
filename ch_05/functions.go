package main

import (
	"errors"
	"fmt"
	"os"
)

// Every Go program starts from a main function. It doesnt take in any parameters or return values.

func main() {
	functions()
}

func functions() {
	fmt.Println(multiply(2, 5))
	fmt.Println(div(5, 2))

	PrintPerson(PrintPersonOpts{
		FirstName: "Gozman",
		LastName:  "Faraday",
		Age:       25,
	})

	fmt.Println(addTo(5))
	fmt.Println(addTo(5, 2, 8, 4, 3, 9, 7))
	slA := []int{1, 2, 4}
	fmt.Println(addTo(5, slA...))
	fmt.Println(addTo(9, []int{2, 5, 7}...))

	result, remainder, err := divAndRemainderOne(5, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

// Other functions can accept parameters and return values

func multiply(num int, denom int) int {
	return num * denom
}

// if two or more consecutive parameters have the same type, you can specify the type for all of them after declaring them.

func div(num, denom float64) float64 {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// Go does not have Named and Optional parameters. But it can be imitated using a struct.

type PrintPersonOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func PrintPerson(opts PrintPersonOpts) {
	fmt.Println("first name:", opts.FirstName)
	fmt.Println("last name:", opts.LastName)
	fmt.Println("age:", opts.Age)
}

//* Variadic Input Parameters and Slices
// Functions in Go can be configured to allow any number of input parameters. We call this variadic Input Parameters & Slices. It is indicated with three dots (...) before the type, and its the last or only parameter in the function. The value of the variadic parameter inside the function is a slice of the specified type.

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out

	// this function can then called like in various ways.
	// addTo(1, 2, 3, 4, 5, 6)
	// a := []int{1, 2, 4}
	// addTo(5, a...)
	// addTo(9, []int{2, 5, 7}...)
}

//* Multiple Return Values
// unlike other languages, a function in Go can return multiple values.
// This shit is amazing. Every return instance in the function must return all the values specified.
// This is commonly used in functions where an error has to be returned if an error occured. If no error occured, `nil` is returned. by convention, the error is always the last value (or only) value returned from a function.

func divAndRemainderOne(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	return num / denom, num % denom, nil
}

// If a function in Go returns multiple values, the returned values must be assigned to multiple variables. Trying to assign all the values to a single variable causes a compile-time error.
// If on eor more of the returned values need to be ignored, you can use the `_` since Go doesnt allow unused variables.

//* Named Return Values
// Go also allows you to name the return values. The names given to the return values, can be used as variables within the function. Named return values are initialized to their zero values when created.
// The name thats used for a named return value remains local to the function.

func divAndRemainderTwo(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}

	result, remainder = num/denom, num%denom
	return result, remainder, err // zero value for an error is nil obvs!
}
