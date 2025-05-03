package main

import (
	"fmt"
)

func main() {
	grade := "B"

	// Evaluate the grade using a switch statement
	// Switch statements provide a cleaner alternative to multiple if-else conditions
	// when comparing a single variable against multiple possible values
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good job!")
	case "C":
		fmt.Println("You can do better.")
	case "D":
		fmt.Println("You need to work harder.")
	case "F":
		fmt.Println("Failed.")
	default:
		// The default case handles any value not explicitly covered above
		fmt.Println("Invalid grade.")
	}
}
