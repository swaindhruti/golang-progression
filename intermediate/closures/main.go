package main

import (
	"fmt"
)

// adder returns a closure that increments a counter and prints its value
// This demonstrates how closures maintain their own state between function calls
func adder() func() {
	i := 0 // This variable is "captured" by the closure
	fmt.Println("Initial value of i:", i)
	return func() {
		i++ // Each time the returned function is called, it increments its own copy of i
		fmt.Println("Added 1 to i:", i)
	}
}

// createSubtractor returns a function that subtracts a given value from an initial sum
func createSubtractor(initialValue int) func(int) int {
	sum := initialValue
	return func(x int) int {
		sum -= x // Each call reduces the captured sum variable by x
		return sum
	}
}

func main() {
	// Example 1: Using the adder closure
	fmt.Println("First sequence:")
	sequence1 := adder() // Create first counter closure

	// Each call to sequence1 increments its internal counter
	sequence1()
	sequence1()
	sequence1()
	sequence1()

	// Example 2: Using a second adder closure
	fmt.Println("\nSecond sequence:")
	sequence2 := adder() // Create second counter closure (independent of the first)

	// Each call to sequence2 increments its own internal counter
	sequence2()
	sequence2()
	sequence2()
	sequence2()

	// Example 3: Using subtractor closure
	fmt.Println("\nSubtractor sequence:")
	subtractor := createSubtractor(100) // Start with 100 and subtract from it

	// Each call returns the result after subtraction
	fmt.Println("After subtracting 1:", subtractor(1))
	fmt.Println("After subtracting 1:", subtractor(1))
	fmt.Println("After subtracting 1:", subtractor(1))
	fmt.Println("After subtracting 1:", subtractor(1))
}
