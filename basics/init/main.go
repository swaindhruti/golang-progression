package main

/*
- In Golang, init functions are special functions that can be declared in any package.
- They are used to perform initialization tasks before the package is used.
- Go executes init functions automatically when the package is initialized.
- This happens before the main function is executed.
*/

func init() {
	// Initialization code can go here
	println("Initializing...")
}

func main() {
	println("Main function executing...")
}
