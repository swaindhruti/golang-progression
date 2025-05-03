package main

import "fmt"

func main() {
	// Using a for loop to iterate over a range of numbers
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// Using range to iterate over a range of numbers
	for i := range 10 {
		fmt.Println("Range Iteration:", i)
	}

	// Using a for loop with a condition
	j := 1
	for j <= 5 {
		fmt.Println("Condition Iteration:", j)
		j++
	}

	// Using a for loop with a break statement
	k := 1
	for {
		if k > 5 {
			break // Exit the loop when k is greater than 5
		}
		fmt.Println("Break Iteration:", k)
		k++
	}

	// Using a for loop with a continue statement
	l := 0
	for l < 10 {
		l++
		if l%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Continue Iteration:", l)
	}

	// Using a for loop with range to iterate over a array
	numbers := []int{4, 8, 7, 6, 5, 4, 3, 2, 1}
	for index, values := range numbers {
		fmt.Printf("Index: %d, Value: %d \n", index, values)
	}
}
