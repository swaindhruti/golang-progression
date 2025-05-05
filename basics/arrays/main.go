package main

import "fmt"

func main() {
	// Arrays in Go are fixed-size collections of elements of the same type.

	var arr [5]int // Declare an array of integers with size 5
	for i := range arr {
		fmt.Printf("Index: %d, ", i)  // Print the index of the array
		fmt.Println("Value:", arr[i]) // Print the value at the index (initially 0)
		arr[i] = i * 2                // Initialize the array with even numbers
	}

	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value) // Print the index of the array
	}

	for _, value := range arr { // Undescored variable _ is called a blank identifier, used to ignore the index
		fmt.Printf("Value: %d\n", value) // Print the value at the index (initially 0)
	}

	fmt.Println("Array:", arr)                  // Print the array
	fmt.Println("Length of array:", len(arr))   // Print the length of the array (number of elements it contains)
	fmt.Println("Capacity of array:", cap(arr)) // Print the capacity of the array (maximum number of elements it can hold without resizing)
	// Note: For arrays, length and capacity are always the same because arrays are fixed-size
	// The difference becomes important when working with slices, where capacity can be larger than length
	fmt.Println("First element:", arr[0])         // Print the first element of the array
	fmt.Println("Last element:", arr[len(arr)-1]) // Print the last element of the array
	fmt.Println("Slice of array:", arr[1:4])      // Print a slice of the array (elements from index 1 to 3)

	originalArray := [5]int{1, 2, 3, 4, 5}                       // Declare and initialize an array
	copiedArray := originalArray                                 // Copy the original array to a new array
	copiedArray[0] = 10                                          // Modify the first element of the copied array
	fmt.Println("Original array after copying:", originalArray)  // Print the original array (unchanged)
	fmt.Println("Copied array after modification:", copiedArray) // Print the copied array (changed)
	// Note: In Go, arrays are value types, so copying an array creates a new copy of the array

	// Comparing arrays
	array1 := [3]int{1, 2, 3}                          // Declare and initialize an array
	array2 := [3]int{1, 2, 3}                          // Declare and initialize another array
	fmt.Println("Are arrays equal?", array1 == array2) // Compare the two arrays for equality
	// Note: Arrays can be compared using the == operator, which checks if all elements are equal

	// Matrix (2D array)
	matrix := [3][3]int{ // Declare and initialize a 2D array (matrix)
		{1, 2, 3}, // First row
		{4, 5, 6}, // Second row
		{7, 8, 9}, // Third row
	}
	fmt.Println("Matrix:", matrix)                  // Print the matrix
	fmt.Println("Element at (1, 2):", matrix[1][2]) // Print the element at row 1, column 2 (5)

	// Using original array in a copied array using pointer and address
	originalArray2 := [5]int{1, 2, 3, 4, 5}                 // Declare and initialize an array
	var pointerArray *[5]int = &originalArray2              // Create a pointer to the original array
	fmt.Println("Pointer to original array:", pointerArray) // Print the pointer to the original array
	fmt.Println("Value at pointer:", *pointerArray)         // Print the value at the pointer (original array)
	// Note: The pointer points to the original array, so any changes made through the pointer will affect the original array

}
