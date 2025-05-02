package main

import "fmt"

var firstname = "Dhruti" // Global variable

func main() {
	fmt.Println("My name is", firstname)
	Printlastname() // Call the function to print the last name
}

// Printlastname prints the author's last name to standard output.
func Printlastname() {
	lastname := "Swain" // Local variable
	fmt.Println("My last name is", lastname)
}
