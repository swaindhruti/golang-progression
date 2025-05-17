package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	   This program demonstrates how os.Exit() works in Go.
	   The key behavior is immediate termination without running
	   any cleanup code or deferred functions.
	*/
	fmt.Println("Starting the main function")

	// Register a deferred function that will never run
	defer fmt.Println("Deferred Statement")

	/*
	   os.Exit() immediately terminates the program with the specified status code:

	   - Status code 0: Indicates successful execution
	   - Status code non-zero (1-255): Indicates an error condition

	   Critical behaviors to understand:
	   - All deferred functions are bypassed (not executed)
	   - No cleanup handlers are run
	   - All goroutines are terminated immediately
	   - The exit status can be observed by the parent process
	*/
	os.Exit(1) // Terminate with error status code 1

	// This line is unreachable code
	fmt.Println("End of main function")
}
