package main

import "fmt"

func main() {

	// Variadic function to add numbers

	numbers := []int{1, 2, 3, 4, 5} // Slice of integers to be passed to the variadic function
	add(numbers...)                 // Calling the variadic function with multiple arguments
}

func add(numbers ...int) { // Variadic function to add numbers
	sum := 0
	for _, number := range numbers { // Looping through the variadic arguments
		sum += number // Adding each number to the sum
	}
	fmt.Println("Sum of numbers is", sum)                    // Printing the sum of the numbers
	fmt.Println("Number of arguments passed:", len(numbers)) // Printing the number of arguments passed

}

// Note: The variadic function can take any number of arguments, including zero arguments. If no arguments are passed, the length of the slice will be zero. Variadic parameters must be the last parameter in the function signature. You can also pass a slice to a variadic function by using the `...` operator, like this: `add(numbers...)`. This will unpack the slice and pass each element as a separate argument to the function.
