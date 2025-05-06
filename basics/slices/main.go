package main

import (
	"fmt"
)

func main() {
	// Slices in Go are dynamic-size collections of elements of the same type.
	// They are more flexible than arrays and can grow or shrink as needed.

	// Declare and initialize a slice of integers
	slice := []int{1, 2, 3, 4, 5} // Slice literal

	// Print the slice
	fmt.Println("Slice:", slice) // Print the slice

	// Append elements to the slice
	slice = append(slice, 6, 7, 8) // Append multiple elements to the slice

	// Print the updated slice
	fmt.Println("Updated Slice:", slice) // Print the updated slice

	// Slice a portion of the original slice
	subSlice := slice[2:5] // Create a sub-slice from index 2 to 4 (inclusive)

	// Print the sub-slice
	fmt.Println("Sub-Slice:", subSlice) // Print the sub-slice

	// Length and capacity of slices
	fmt.Println("Length of Slice:", len(slice))   // Print the length of the slice (number of elements it contains)
	fmt.Println("Capacity of Slice:", cap(slice)) // Print the capacity of the slice (maximum number of elements it can hold without resizing)

	// Copying slices
	copiedSlice := make([]int, len(slice)) // Create a new slice with the same length as the original slice
	copy(copiedSlice, slice)               // Copy elements from the original slice to the new slice

	// Print the copied slice
	fmt.Println("Copied Slice:", copiedSlice) // Print the copied slice

	// Comparing slices (not directly possible in Go)
	isEqual := true        // Initialize a boolean variable to true
	for i := range slice { // Iterate over each element in the original slice
		if slice[i] != copiedSlice[i] { // Compare elements at index i in both slices
			isEqual = false // If any element is not equal, set isEqual to false
			break           // Break out of the loop if not equal
		}
	}
	fmt.Println("Are slices equal?", isEqual) // Print whether slices are equal or not

	// Using nil slices
	var nilSlice []int      // Declare a nil slice (no memory allocated)
	if len(nilSlice) == 0 { // Check if the nilSlice is empty (a better way to check)
		fmt.Println("Nil Slice is empty")
	}
}

// Note: Slices are reference types, so when you pass a slice to a function, it passes a reference to the underlying array, not a copy of the array. This means that changes made to the slice inside the function will affect the original slice outside the function.
// Slices can be resized, and their capacity can be larger than their length. When you append elements to a slice and it exceeds its capacity, Go automatically allocates a new array with double the capacity and copies the elements to the new array.
// This is known as slice growth. The new slice will have a new underlying array, and the original slice will remain unchanged.
// Slices can be created using the built-in make function, which allows you to specify the length and capacity of the slice. For example, make([]int, 5, 10) creates a slice of integers with length 5 and capacity 10. The length is the number of elements currently in the slice, while the capacity is the maximum number of elements it can hold without resizing.
