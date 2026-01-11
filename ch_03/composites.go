package main

import (
	"fmt"
	"slices"
)

func main() {
	// Composite Types — Arrays, Slices, and Maps?

	array()
	slice()
}

func array() {
	//* ARRAYS
	// Arrays in Go are too rigid. They're rarely used.

	// To declare an array, you can specify the size of the array and the type of the elements in the array.
	// Once an array has been declared, the size and the type cannot be changed. Like I said, it's quite rigid.
	// When declared this way, the elements of the array, will have the zero value of the assigned type.

	var a1 [3]int // [0, 0, 0]

	// The value for the arrays can also be specified when declaring them.

	a2 := [5]int{1, 2, 3, 4, 5} // [1, 2, 3, 4, 5]

	// If some elements need to have the zero value of the assigned type, the nonzero elements can specified.

	a3 := [5]int{1, 2: 24, 4: 100} // [1, 0, 24, 0, 100]

	// when declaring an array with array literals, the size of the array can be replaced with `...`. The size is implied from the array literal.

	a4 := [...]float64{1.2, 3.4, 5.6} // [1.2, 3.4, 5.6]

	// Two arrays are the same if they're of the same type, size and have the same values.

	// To create multidimensional arrays, we can do:

	var mA1 [3][4]int // An array of 3 arrays of 4 strings.

	// Arrays in Go are read using the bracket notation.

	a5 := mA1[2]

	// Arrays in Go have some limitations. The size of the array is considered part of the type of the array. This has certain consequences. You can't use variables to specify the size of an array. An array of [3]int is diff from an array of [4]int. They cannot even be type converted into one another. You also cannot assign arrays of diff sizes to the same variable.

	fmt.Println(a1, a2, a3, a4, mA1, a5)
}

func slice() {
	//* SLICES
	// Slices are different from arrays because the size isn't a part of it's type. It can also grow/shrink as needed.

	s1 := []int{1, 2, 3} // [1, 2, 3]

	// Specific indices can also be specified as nonzero just like arrays.

	s2 := []float64{1.2, 4: 22.5, 10, 8: 9.81} // [1.2, 0, 0, 0, 22.5, 10, 0, 0, 9.81]

	// We can also make multidimensional slices just like in arrays

	var mS1 [][]int

	// Elements in a slice is also read or written using the bracket notation.
	// Just as in arrays, you can't read or write past the end of the slice or use a negative index.

	s2[1] = 5.8 // [1.2, 5.8, 0, 0, 22.5, 10, 0, 0, 9.81]

	// Slices are also diff from Arrays in certain ways

	// The zero value of a slice declared without a literal (empty slice) is `nil`.

	var eS []string

	// Two slices cannot be compared ot each other. Doing it will throw an error.
	// A slice can only be compared with `nil`.
	// A slice with a value of `nil` means the slice is empty.

	// We can compare slices using the `slices` package from the standard library.
	// If the slices both have the same value, type and size, the `slices.Equal` func will return true.

	s3 := []int{1, 2, 3}
	sIE := slices.Equal(s1, s3)

	// The length of a slice can gotten using the func `len()`. Passing a `nil` sloce to len returns 0.

	var s4 []int

	fmt.Println(s1, s2, mS1, eS == nil, sIE, len(s3), len(s4))

	//* append
	// We can append elements of a valid type to a slice of that type.

	s4 = append(s4, 10)

	// we can append more than one element at a time

	s4 = append(s4, 20, 30)

	// we can also append elements from another slice by using `...`.

	s4 = append(s4, s3...)

	//* Capacity
	// Each element in a slice is assigned to consecutive memory locations. Every slice has a capacity. The capacity of a slice is the number of consecutive memory locations reserved.
	// We can check the capacity of a slice using `cap()`. The capacity of an array is always equal to the length of that array. The capacity of a slice can be greater than its length.

	fmt.Println("s4 value, length and capacity:", s4, len(s4), cap(s4))

	//* make()
	// There might be times when you know the number of things you want to put into a slice. While, its a good thing that slices can grow, its far more efficient to create them with the correct initial capacity.
	// We can do that using `make()`

	sLC1 := make([]int, 5) // this creates a slice of len and cap of 5. The zero value isn't `nil`.

	sLC2 := make([]float64, 4, 10) // a slice of length 4, and capacity 10.

	// NOTE: if you make a slice using `make`, appending elements will append them to the end of the already created slice. Appending always increase the length.

	fmt.Println(sLC1, sLC2)

	//* Emptying a Slice
	// We can turn all the elements of a slice to the corresponding zero value using the `clear()` function.

	clear(s4)
	fmt.Println(s4)

	//* Slicing Slices
	// We can slice out Slices from other slices, using the notation slice[a:b], where a is the index of the first element of the new slice and b is the index + 1 of the last element of the new slice.
	// a or/and b can be left out. When a is removed [:b], the new slice is taken from the first element to the  element just before the element with index of b. When b is removed [a:], the new slice is taken from the element of index a to the last element if the old slice. When both are removed [:], the new slice is a copy of the old slice.

	s5 := []string{"a", "b", "c", "d"}

	s5S1 := s5[1:3]
	s5S2 := s5[:2]
	s5S3 := s5[1:]
	s5S4 := s5[:]

	fmt.Println(s5, s5S1, s5S2, s5S3, s5S4)

	// When you create a slice from another slice, a copy isn't being made. Rather, you get two variables that are sharing the same memory. This means that a change to an original slice, will affect all slices that share the same memory (was sliced out from that original slice).

	s6X := []string{"a", "b", "c", "d"}
	s6Y := s6X[:2] // ["a", "b"]
	s6Z := s6X[1:] // ["b", "c", "d"]
	s6X[1] = "y"
	s6Y[0] = "x"
	s6Z[1] = "z"

	fmt.Println("X:", s6X) // ["x", "y", "z", "d"]
	fmt.Println("Y:", s6Y) // ["x", "y"]
	fmt.Println("Z:", s6Z) // ["y", "z", "d"]

  //* copy
}
