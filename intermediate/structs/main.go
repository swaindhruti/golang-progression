package main

import (
	"fmt"
)

// Person represents an individual with personal information
type Person struct {
	firstName     string
	lastName      string
	Age           int
	Address       Address // Embedded struct for address details
	PhoneHomeCell         // Embedded struct that promotes fields to Person level
}

// Address represents a physical location
type Address struct {
	City    string
	State   string
	Country string
}

// PhoneHomeCell contains home and cell phone numbers
// When embedded, its fields are directly accessible from the parent struct
type PhoneHomeCell struct {
	Home string
	Cell string
}

func main() {
	// Create a Person instance with nested structs
	person := Person{
		firstName:     "Alice",
		lastName:      "Wesly",
		Age:           30,
		Address:       Address{City: "New York", State: "NY", Country: "USA"},
		PhoneHomeCell: PhoneHomeCell{Home: "123-456-7890", Cell: "987-654-3210"},
	}

	// Access and print Person fields
	fmt.Println("Person struct:", person)
	fmt.Println("Name:", person.firstName)
	fmt.Println("Last Name:", person.lastName)

	// Access nested Address struct fields
	fmt.Println("City:", person.Address.City)
	fmt.Println("State:", person.Address.State)
	fmt.Println("Country:", person.Address.Country)

	// Access promoted fields from PhoneHomeCell (directly available on Person)
	fmt.Println("Home:", person.Home) // Promoted field access
	fmt.Println("Cell:", person.Cell) // Promoted field access

	// Demonstrate method usage
	fmt.Println("Age Before Increment:", person.Age)
	person.incrementAge()                           // Call method with pointer receiver
	fmt.Println("Age After Increment:", person.Age) // Value was modified
	fmt.Println("Full Name:", person.fullName())    // Call method with value receiver

	// Demonstrate struct comparison
	person2 := Person{firstName: "John", lastName: "Doe", Age: 40}
	fmt.Println("Are person and person2 the same?", person == person2) // Compare structs

	// Anonymous struct example
	user := struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  25,
	}
	fmt.Println("User struct:", user)
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
}

// fullName is a method with a value receiver that returns the person's full name
func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// incrementAge is a method with a pointer receiver that modifies the person's age
// Using a pointer receiver allows the method to modify the original Person
func (p *Person) incrementAge() {
	p.Age++ // Increment the Age field by 1
}
