package main

import (
	"fmt"
	foo "net/http" // Importing net/http package with alias 'foo'
)

func main() {
	fmt.Println("Hello, World!") // Print a greeting message

	// Make an HTTP GET request to the JSONPlaceholder API
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println("Error:", err) // Print any errors that occurred during the request
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed when the function returns

	fmt.Println("Response Status:", resp.Status) // Print the HTTP status of the response
}
