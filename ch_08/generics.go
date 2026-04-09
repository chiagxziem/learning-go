package main

import "fmt"

func main() {
	Generics()
}

func Generics() {
	//* Generics
	// Generics is like a type params. They allow us to write functions or types that work with various types without sacrificing type safety or resorting to `any`.

	var intStack Stack[int]
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	// intStack.Push("wtf") // this will panic because intStack was initialized using int as T.
	v, ok := intStack.Pop()
	fmt.Println(v, ok)
	fmt.Println(intStack.Contains(10))
	fmt.Println(intStack.Contains(5))

  //* Generic Functions
}

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		// to return the zero value of whatever type is passed as T, just declare a variable with
		// type T and no assingment. this will always return the zero value of that type.
		var zero T
		return zero, false
	}

	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:(len(s.vals) - 1)]
	return top, true
}
func (s *Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false

	// can all be replaced by `slices.Contains(s.vals, val)`
}
