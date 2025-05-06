package main

import (
	"fmt"
	"maps"
)

func main() {
	// Maps in Go are unordered collections of key-value pairs.
	// They are similar to dictionaries in Python or hash tables in other languages.
	// Maps are reference types, meaning they are passed by reference and can be modified in place.

	myMap := make(map[string]int) // Create a new map with string keys and int values

	fmt.Println("Map:", myMap) // Print the map (initially empty)

	myMap["one"] = 1   // Add a key-value pair to the map
	myMap["two"] = 2   // Add another key-value pair to the map
	myMap["three"] = 3 // Add another key-value pair to the map
	myMap["four"] = 4  // Add another key-value pair to the map
	myMap["five"] = 5  // Add another key-value pair to the map

	fmt.Println("Updated Map:", myMap)                // Print the map
	fmt.Println("Length of map:", len(myMap))         // Print the length of the map (number of key-value pairs it contains)
	fmt.Println("Value for key 'two':", myMap["two"]) // Print the value for the key "two"
	fmt.Println("Value for key 'six':", myMap["six"]) // Print the value for the key "six" (not present in the map)

	// Printing the map using a for loop
	for key, value := range myMap { // Loop through the map
		fmt.Printf("Key: %s, Value: %d\n", key, value) // Print the key and value
	}

	myMap2 := map[string]int{ // Declare and initialize a map with key-value pairs
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	fmt.Println("Map 2:", myMap2) // Print the second map

	// Compare the two maps for equality
	if maps.Equal(myMap, myMap2) {
		fmt.Println("Maps are equal") // Print if the maps are equal
	} else {
		fmt.Println("Maps are not equal") // Print if the maps are not equal
	}

	myMap2["three"] = 21                                      // Update the value for the key "three" in the second map
	fmt.Println("Updated Map 2:", myMap2)                     // Print the updated second map
	fmt.Println("Accesing non existing key: ", myMap2["six"]) // Note: Accessing a non-existing key in a map returns the zero value for the value type (0 for int, "" for string, etc.)

	value, exists := myMap2["three"]                // Accessing a key-value pair in the map
	_, ok := myMap2["six"]                          // Check if the key "six" exists in the map, Note: Conventionally ok is used for the boolean variable indicating existence and the underscore _ is used to ignore the value returned by the map access
	fmt.Println("Key 'three' exists:", ok)          // Print the existence boolean for the key "three"
	fmt.Println("Value:", value, "Exists:", exists) // Print the value and existence boolean

	delete(myMap2, "two")                            // Delete the key "two" from the map
	fmt.Println("Deleted key 'two':", myMap2)        // Print the map after deletion
	fmt.Println("Deleted key 'six':", myMap2["six"]) // Print the value for the key "six" (not present in the map)

	// Clearing a map
	clear(myMap2)                         // Clear the map by reinitializing it to an empty map
	fmt.Println("Cleared Map 2:", myMap2) // Print the cleared map
	// Note: Maps can be cleared by reinitializing them to an empty map or by deleting all key-value pairs using the delete function

	var myMap3 map[string]int                    // Declare a nil map (no memory allocated)
	fmt.Println("Is myMap3 nil?", myMap3 == nil) // Check if the map is nil
	// Note: A nil map is a map that has not been initialized. It cannot be used until it is initialized using make or a map literal.
	// Note: Maps are reference types, so they can be passed to functions and modified in place.

	// Nested maps
	nestedMap := map[string]map[string]int{ // Declare a nested map (map of maps)
		"outerKey1": { // Outer key "outerKey1" with an inner map
			"innerKey1": 1,
			"innerKey2": 2,
		},
		"outerKey2": { // Outer key "outerKey2" with an inner map
			"innerKey3": 3,
			"innerKey4": 4,
		},
	}

	fmt.Println("Nested Map:", nestedMap)                                               // Print the nested map
	fmt.Println("Value for outerKey1, innerKey1:", nestedMap["outerKey1"]["innerKey1"]) // Print the value for the outer key "outerKey1" and inner key "innerKey1"

}
