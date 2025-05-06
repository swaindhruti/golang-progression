package main

import (
	"fmt"
)

func main() {
	// Define a string to iterate over
	message := "Hello, World!"

	// Using range to iterate over the characters in a string
	// The range keyword provides both the index and the rune (Unicode code point)
	// Note: In Go, strings are UTF-8 encoded, so each character might use 1-4 bytes
	for i, char := range message {
		// Print the index, the character as a rune, and its Unicode decimal value
		// %d = decimal integer, %c = character representation
		fmt.Printf("Character at %d: %c (Unicode: %d)\n", i, char, char)
	}

	fmt.Println("\nIgnoring the index:")

	// When you don't need the index, you can use the blank identifier (_)
	// This is more memory efficient as the index value isn't allocated
	for _, char := range message {
		fmt.Printf("Character: %c (Unicode: %d)\n", char, char)
	}

	// Using the blank identifier (_) is a Go convention that:
	// 1. Makes code more readable by showing you're intentionally ignoring a value
	// 2. Prevents the compiler from allocating memory for unused variables
	// 3. Avoids the "declared but not used" compilation error
}
