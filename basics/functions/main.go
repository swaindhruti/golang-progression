package main

import (
	"fmt"
)

func main() {

	// Calling a function to add two numbers
	fmt.Println("Sum of 2 and 3 is", add(2, 3))

	greet := func() { // Anonymous function to greet the user
		fmt.Println("Hello, User!")
		fmt.Println("This is an anonymous function")
	}

	greet() // Calling the anonymous function

	operation := add          // Assigning the add function to a variable
	result := operation(5, 7) // Assigning the add function to a variable and calling it
	fmt.Println("Sum of 5 and 7 is", result)

	// Using a function as an argument
	result2 := applyOperation(10, 20, add) // Passing the add function as an argument
	fmt.Println("Sum of 10 and 20 is", result2)

	multiplierFunc := makeMultiplier(2)              // Creating a multiplier function with a factor of 2
	result3 := multiplierFunc(5)                     // Assigning the multiplier function to a variable
	fmt.Println("Multiplying 5 by 2 gives", result3) // Calling the multiplier function with 5

	coordnateX, coordinateY := getCoordinates()              // Calling a function that returns multiple values
	fmt.Println("Coordinates are:", coordnateX, coordinateY) // Printing the coordinates returned by the function

	result4Q, result4R := divide(10, 3) // Calling a function that returns multiple values with named return values
	fmt.Println("Quotient :", result4Q, "Remainder :", result4R)

}

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}

// Function that takes another function as an argument
func applyOperation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Function that returns another function
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Function returnig multiple values
func getCoordinates() (int, int) {
	return 10, 20
}

// Function returning multiple values with named return values
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// Note: In Go, functions are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions. This allows for a high degree of flexibility and modularity in code design.
// Anonymous functions are functions that are defined without a name and can be used for short-lived tasks or callbacks. They can also capture variables from their surrounding context, which is useful for closures.
// In the example above, we defined a function `add` that takes two integers and returns their sum. We also created an anonymous function `greet` that prints a greeting message. Finally, we assigned the `add` function to a variable `operation` and called it with different arguments.
// This demonstrates the flexibility of functions in Go and how they can be used in various ways to create modular and reusable code.
// The example also shows how to use anonymous functions for short-lived tasks or callbacks, and how to capture variables from their surrounding context using closures.
