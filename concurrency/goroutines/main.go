package main

import (
	"fmt"
)

func greet(name string) {
	for range 3 {
		fmt.Println("Hello", name)
		// time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// "Alice" starts at the same time "Bob" starts too because it's a goroutine and
	// therefore is non-blocking. "Alice" runs concurrently.

	//! PROBLEM
	// Main might exit and the program terminate before "Alice" finishes.
	// Comment out the `time.Sleep` line and run the program to notice that "Alice"
	// doesn't get printed at all. We need coordination.

	// This is fixed using wait groups.

	go greet("Alice")
	greet("Bob")
}
