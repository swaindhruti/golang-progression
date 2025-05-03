package main

import "fmt"

func main() {
	// This program demonstrates the use of a while loop in Go.
	// In golang there is no while loop, but we can use for loop as a while loop.
	// The for loop can be used to create a while loop by omitting the initialization and post statements.
	// The loop will continue until the condition is false.

	var i int = 0 // Initialize a variable to use as a counter

	for {
		if i > 5 { // Check the condition
			break // Exit the loop if the condition is met
		}
		fmt.Println("Iteration:", i) // Print the current value of i
		i++                          // Increment i by 1
	}

}
