package main

import "fmt"

func main() {

	// Defer is used to ensure that a function call is performed later in the program even if the function returns an error.
	// It is often used for cleanup actions, such as closing files or releasing resources.

	process(10)

}

func process(i int) { // The deferred function calls are executed in LIFO order (Last In First Out).
	defer fmt.Println("The value of i is: ", i)             // Executed 10th
	fmt.Println("Processing-1.")                            // Executed 1st
	defer fmt.Println("Second deferred statement executed") // Executed 9th
	fmt.Println("Processing-2.")                            // Executed 2nd
	defer fmt.Println("Third deferred statement executed")  // Executed 8th
	i++                                                     // Executed 3rd
	fmt.Println("Normal execution statement")               // Executed 4th
	fmt.Println("The actual value of i is: ", i)            // Executed 5th
	defer fmt.Println("Fourth deferred statement executed") // Executed 7th
	defer fmt.Println("The value of i is: ", i)             // Executed 6th
}

// Note: The defered statement will be exectured immediately in the function but it will be executed at the end of the function execution.
// Practical Use Cases: Resource Cleanup, Unlock Mutexes, Loging and Tracking
// Best Practices: Keed defered actions short, Understand Evaluation Timings, Avoid Complex Control Flow
// Commonly Used to ensure that resources like files or te DB connections are closed after they are opened.
