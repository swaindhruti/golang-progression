package main

import (
	"fmt"
)

func main() {
	for i := range 20 {
		// This program demonstrates the use of if-else statements in Go.
		// The if-else statement is used to execute different blocks of code based on a condition.
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}

	// The if-else statement can also be used to check multiple conditions using else if.
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// Nested if-else statements can be used to check multiple conditions.
	age := 25
	if age >= 18 {
		fmt.Println("You are an adult.")
		if age >= 21 {
			fmt.Println("You can drink alcohol.")
		} else {
			fmt.Println("You cannot drink alcohol.")
		}
	} else {
		fmt.Println("You are a minor.")
	}
	// The if statement can also be used without an else block.
}
