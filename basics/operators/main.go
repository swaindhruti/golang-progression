package main

import (
	"fmt"
)

func main() {

	// Logical Operators
	a := true
	b := false

	fmt.Println("Logical AND:", a && b) // false
	fmt.Println("Logical OR:", a || b)  // true
	fmt.Println("Logical NOT:", !a)     // false

	// Bitwise Operators
	x := 5 // 0101 in binary
	y := 3 // 0011 in binary

	fmt.Println("Bitwise AND:", x&y)  // 0001 in binary (1 in decimal)
	fmt.Println("Bitwise OR:", x|y)   // 0111 in binary (7 in decimal)
	fmt.Println("Bitwise XOR:", x^y)  // 0110 in binary (6 in decimal)
	fmt.Println("Bitwise NOT:", ^x)   // 1010 in binary (10 in decimal)
	fmt.Println("Left Shift:", x<<1)  // 1010 in binary (10 in decimal)
	fmt.Println("Right Shift:", x>>1) // 0010 in binary (2 in decimal)

	// Comparison Operators
	fmt.Println("Equal to:", x == y)                 // false
	fmt.Println("Not Equal to:", x != y)             // true
	fmt.Println("Greater than:", x > y)              // true
	fmt.Println("Less than:", x < y)                 // false
	fmt.Println("Greater than or equal to:", x >= y) // true
	fmt.Println("Less than or equal to:", x <= y)    // false

}
