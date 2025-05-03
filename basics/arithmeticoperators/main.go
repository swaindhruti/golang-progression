package main

import "fmt"

func main() {
	var a, b int = 10, 20 // Initialize two integers
	var result int        // Declare a variable to store the result

	result = a + b // Addition
	fmt.Println("Addition:", result)

	result = a - b // Subtraction
	fmt.Println("Subtraction:", result)

	result = a * b // Multiplication
	fmt.Println("Multiplication:", result)

	result = a / b // Division
	fmt.Println("Division:", result)

	result = a % b // Modulus
	fmt.Println("Modulus:", result)

	var pi float64 = 22 / 7 // Integer division
	fmt.Println("Integer Division with constant:", pi)

	pi = 22 / 7.0 // Floating-point division
	fmt.Println("Floating-point Division:", pi)

	// Overflow example with signed integers
	// Note: In Go, integer overflow does not panic; it wraps around.
	var maxInt int64 = 9223372036854775807 // Maximum value for int64
	maxInt = maxInt + 1                    // This will cause an overflow
	fmt.Println("Overflow example:", maxInt)

	// Overflow example with unsigned integers
	var maxUint uint64 = 18446744073709551615 // Maximum value for uint64
	maxUint = maxUint + 1                     // This will cause an overflow
	fmt.Println("Overflow example with unsigned integer:", maxUint)

	// Underflow example
	var smallFloat float64 = -1.7976931348623157e+308 // Minimum value for float64
	smallFloat = smallFloat - 1                       // This will cause an underflow
	fmt.Println("Underflow example:", smallFloat)

	// Bitwise operators
	var x, y int = 5, 3               // Initialize two integers
	fmt.Println("Bitwise AND:", x&y)  // Bitwise AND
	fmt.Println("Bitwise OR:", x|y)   // Bitwise OR
	fmt.Println("Bitwise XOR:", x^y)  // Bitwise XOR
	fmt.Println("Bitwise NOT:", ^x)   // Bitwise NOT
	fmt.Println("Left Shift:", x<<1)  // Left Shift
	fmt.Println("Right Shift:", x>>1) // Right Shift
}
