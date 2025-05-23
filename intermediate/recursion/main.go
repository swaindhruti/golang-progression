package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// Calculate and print the 4th Fibonacci number
	result := fibonacci(4)
	fmt.Println(result)
}

/* fibonacci calculates the nth number in the Fibonacci sequence
The sequence starts with 0, 1, 1, 2, 3, 5, 8, ...
where each number is the sum of the two preceding ones
*/

func fibonacci(n int) int {
	// Base cases
	if n <= 0 {
		return 0 // Return 0 for invalid inputs
	}

	switch n {
	case 1:
		return 0 // First number in sequence is 0
	case 2:
		return 1 // Second number in sequence is 1
	default:
		// For any other position, sum the previous two numbers
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
