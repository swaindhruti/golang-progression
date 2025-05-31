package main

import (
	"fmt"
)

func main() {

	// Printing Functions
	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "Dhruti"
	age := 25
	fmt.Printf("Name: %s, Age : %d\n", name, age)
	fmt.Printf("Binary: %b, Hexadecimal : %X\n", age, age)

	// Formatting Functions
	s := fmt.Sprint("Hello ", "World!", 123, 456) // Sprint returns a formatted string without a newline at the end and spaces between arguments
	fmt.Print(s)

	s = fmt.Sprintln("Hello ", "World!", 123, 456) // Sprintln returns a formatted string with a newline at the end and spaces between arguments
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age : %d\n", name, age) // Sprintf returns a formatted string
	// Sprintf formats according to a format specifier and returns the resulting string and no newline at the end
	fmt.Print(sf)

	// Scanning Functions
	var str string
	var num int

	fmt.Print("Enter a string and a number: ")
	fmt.Scan(&str, &num) // Scan reads input from standard input and stores it in the provided variables
	fmt.Printf("You entered: %s and %d\n", str, num)

	// Clear the input buffer by reading and discarding the newline character
	var discard string
	fmt.Scanln(&discard)

	var str2 string
	var num2 int

	fmt.Print("Enter a string and a number: ")
	fmt.Scanln(&str2, &num2) // Scanln reads a line of input and stores values in provided variables
	fmt.Printf("You entered: %s and %d\n", str2, num2)

	// Error Handling
	fmt.Print("Enter your age: ")
	var ageInput int
	fmt.Scan(&ageInput)      // Scan reads input and returns the number of items successfully scanned and an error if any
	op := checkAge(ageInput) // Call the checkAge function to validate the age
	if op != nil {
		fmt.Println(op) // If checkAge returns an error, print the error message
		return
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is not valid, must be 18 or older", age) // fmt.Errorf creates an error with a formatted message
	}
	return nil
}
