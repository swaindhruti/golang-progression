package main

import "fmt"

var firstname = "Dhruti" // Global variable (string)
var age = 30             // Global variable (int)
var height = 5.8         // Global variable (float64)
var isStudent = true     // Global variable (bool)

func main() {
	fmt.Println("My name is", firstname)
	fmt.Println("My age is", age)
	fmt.Println("My height is", height, "feet")
	fmt.Println("Student status:", isStudent)

	// Array example
	var skills = [3]string{"Go", "Python", "JavaScript"}
	fmt.Println("My skills:", skills)

	// Slice example
	hobbies := []string{"Reading", "Coding", "Hiking"}
	fmt.Println("My hobbies:", hobbies)

	// Map example
	contactInfo := map[string]string{
		"email": "dhruti@example.com",
		"phone": "123-456-7890",
	}
	fmt.Println("Contact info:", contactInfo)

	// HTML template example
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<h1>Hello {{.Name}}</h1>
	<p>Welcome to Go programming!</p>
</body>
</html>
`
	fmt.Println("HTML Template example:")
	fmt.Println(htmlTemplate)

	Printlastname() // Call the function to print the last name
}

// Printlastname prints the author's last name to standard output.
func Printlastname() {
	lastname := "Swain" // Local variable
	fmt.Println("My last name is", lastname)
}
