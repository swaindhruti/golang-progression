package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of terms in the Fibonacci series: ")
	fmt.Scan(&n)

	var a, b int = 0, 1 // Initialize the first two terms of the series
	if n < 0 {
		fmt.Println("Please enter a positive integer")
		return
	} else {
		switch n {
		case 0:
			fmt.Println("Fibonacci series: []")
		case 1:
			fmt.Println("Fibonacci series: [0]")
		default:
			fmt.Print("Fibonacci series: [")
			for i := 0; i < n; i++ {
				if i == n-1 {
					fmt.Print(a)
				} else {
					fmt.Print(a, ", ")
				}
				nextTerm := a + b // Calculate the next term
				a = b             // Update a to the next term
				b = nextTerm      // Update b to the next term
			}
			fmt.Println("]") // Print a closing bracket after the series
		}
	}

}
