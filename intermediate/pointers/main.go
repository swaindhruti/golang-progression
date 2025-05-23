package main

import (
	"fmt"
)

func main() {

	/*
	 * Pointer Demonstration
	 * A pointer is a variable that stores the memory address of another variable.
	 * We use '&' to get the address of a variable and '*' to dereference a pointer
	 * (access the value stored at that address).
	 */

	// Initialize an integer variable
	var a int = 10

	// Declare a pointer to an integer
	var ptr *int

	// Assign the address of 'a' to the pointer
	ptr = &a

	// Print the value of 'a'
	fmt.Println("Value of a:", a)

	// Print the memory address stored in the pointer
	fmt.Println("Address stored in ptr:", ptr)

	// Print the value at the address that ptr points to (dereferencing)
	fmt.Println("Value at the address ptr points to:", *ptr)
}
