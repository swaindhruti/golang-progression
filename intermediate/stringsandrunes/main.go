// Package main demonstrates string and rune handling in Go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// SECTION: Different string representations
	// String with escape sequences
	message := "hello,\nworld!\n"  // \n creates a newline
	message1 := "hello, \tworld\n" // \t creates a tab
	message2 := "hello, \rgo\n"    // \r creates a carriage return

	// Raw string literal (backticks) preserves literal characters
	rawMessage := `hello\nworld` // \n is treated as two characters, not a newline

	// Print each string to see how they render
	fmt.Println("--- String Output ---")
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	// SECTION: String lengths
	// In Go, len() returns the number of bytes in a string
	fmt.Println("\n--- String Lengths (in bytes) ---")
	fmt.Println("Length of message:", len(message))       // 14 bytes (includes newlines as single bytes)
	fmt.Println("Length of message1:", len(message1))     // 14 bytes
	fmt.Println("Length of message2:", len(message2))     // 11 bytes
	fmt.Println("Length of rawMessage:", len(rawMessage)) // 11 bytes (\\ and n counted separately)

	// SECTION: Basic string operations
	fmt.Println("\n--- Basic String Operations ---")
	// Access first byte of string (returns ASCII/UTF-8 value)
	fmt.Println("First byte of message:", message[0]) // 104 (ASCII for 'h')

	// String concatenation
	greeting := "hello "
	name := "dhruti"
	fmt.Println("Concatenation:", greeting+name) // "hello dhruti"

	// SECTION: String comparisons
	fmt.Println("\n--- String Comparisons ---")
	str1 := "Apple"
	str2 := "Banana"
	str3 := "App"
	str4 := "app"

	// Strings are compared lexicographically (byte by byte)
	fmt.Println("Apple < Banana:", str1 < str2) // true: 'A' comes before 'B'
	fmt.Println("Apple > App:", str1 > str3)    // true: longer string with matching prefix
	fmt.Println("Apple > app:", str1 > str4)    // false: uppercase 'A' < lowercase 'a' in ASCII

	// SECTION: Unicode and runes
	fmt.Println("\n--- Unicode and Runes ---")
	// Iterating over string by rune (Unicode code point)
	fmt.Println("Rune-by-rune iteration of message:")
	for index, char := range message {
		fmt.Printf("Position %d: %c (decimal: %d, hex: %x)\n",
			index, char, char, char)
	}

	// Counting runes (characters) vs bytes
	fmt.Printf("Rune count: %d, Byte count: %d\n",
		utf8.RuneCountInString(message), len(message))

	// SECTION: Rune examples
	fmt.Println("\n--- Rune Examples ---")
	var ch rune = 'a'   // rune is an alias for int32
	japaneseChar := 'け' // Unicode character (Hiragana "ke")

	// Convert rune to string
	convertedString := string(ch)

	fmt.Println("Rune value of 'a':", ch)                // 97 (decimal value)
	fmt.Printf("Rune as character: %c\n", ch)            // 'a' (character representation)
	fmt.Printf("Type of rune: %T\n", ch)                 // int32
	fmt.Printf("Japanese character: %c\n", japaneseChar) // け
	fmt.Println("Rune converted to string:", convertedString)
	fmt.Printf("Type of converted string: %T\n", convertedString)
}
