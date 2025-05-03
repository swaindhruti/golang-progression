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

	// Using fallthrough to execute multiple cases
	grade2 := "A"

	// The fallthrough statement allows the execution to continue to the next case
	// even if the current case matches. This is useful for executing common code
	// across multiple cases without repeating it.
	// Note: fallthrough only works for the next case, not for any case further down
	switch grade2 {
	case "A":
		fmt.Println("Excellent!")
		fallthrough // This will execute the next case as well
	case "B":
		fmt.Println("Good job!")
	case "C":
		fmt.Println("You can do better.")
	case "D":
		fmt.Println("You need to work harder.")
	case "F":
		fmt.Println("Failed.")
	default:
		fmt.Println("Invalid grade.")
	}

	// Checking multiple conditions in a single case
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday": // Multiple values can be checked in a single case
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Using expression in switch cases
	number := 10

	// The switch statement can also evaluate expressions, not just values
	switch {
	case number < 0:
		fmt.Println("Negative number")
	case number >= 0 && number <= 10: // Using a range of values in the condition
		fmt.Println("Number is between 0 and 10")
	case number > 10 && number <= 20:
		fmt.Println("Number is between 11 and 20")
	default:
		fmt.Println("Number is greater than 20")
	}

	var x any = 42 // Using the empty interface to hold any type. The empty interface is a type that can hold any value.
	checkType(x)   // Call the function to check the type of x
}

func checkType(x any) { // any is an empty interface that can hold any type of value
	// The switch statement can also be used to check the type of a variable
	switch x.(type) {
	case int:
		fmt.Println("x is an integer")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a boolean")
	default:
		fmt.Println("Unknown type")
	}
}
