package main

/*
- In Golang, init functions are special functions that can be used for initialization tasks
- They execute automatically when the package is initialized, before main
- They run in the order they are defined in the source code
- They take no parameters and return no values
*/

// init functions can be grouped by purpose
func init() {
	println("Initializing...1 - First phase")
}

// Separate init functions with clear comments about their purpose
func init() {
	println("Initializing...2 - Configuration")
}

// Each init function can handle a different aspect of initialization
func init() {
	println("Initializing...3 - Resource setup")
}

// Final initialization steps
func init() {
	println("Initializing...4 - Ready to start")
}

func main() {
	println("Main function executing...")
}
