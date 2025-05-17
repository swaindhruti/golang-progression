package main

import "fmt"

func main() {
	// Example of normal execution
	process(10)

	// Example that will trigger panic
	process(-10)

}

func process(input int) {
	// Deferred statements are executed in LIFO order (Last In, First Out)
	// when the function returns, whether normally or due to panic
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	// Check if input is valid
	if input < 0 {
		fmt.Println("Before Panic")
		// Panic immediately stops normal execution and unwinds the stack,
		// running any deferred functions before terminating the program
		panic("Input must be non-negative")
	}

	fmt.Println("Processing Input:", input)
}
