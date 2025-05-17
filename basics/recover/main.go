package main

import "fmt"

func main() {
	// Demonstrate panic recovery mechanism
	process()

	// This line will execute because the panic was recovered
	fmt.Println("Return from process")
}

func process() {
	// Defer a recovery function that will execute when process() exits
	// This must be registered before any potential panic occurs
	defer func() {
		// recover() returns the value passed to panic() or nil if no panic occurred
		if r := recover(); r != nil { // This line checks if recover caught a panic  and r will not be nil.
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")

	// This causes immediate termination of the current function
	// and begins unwinding the stack, executing deferred functions
	panic("Something went wrong")

}
