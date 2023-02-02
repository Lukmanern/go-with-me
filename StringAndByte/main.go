package main

import "fmt"

func main() {
	// Call the compare function
	compare()
	// Call the byte_to_string function
	byte_to_string()
}

// Compare function to compare the 
// difference between string and byte
func compare() {
	// Print the string representation of "a"
	fmt.Println(`string => "a" :`, "a")
	// Print the byte representation of 'a'
	fmt.Println(`byte => 'a' :`, 'a')
}

// ByteToString function to
// convert a byte to a string
func byte_to_string() {
	// Declare and initialize a byte 
	// variable with the value 'b'
	var b byte = 'b'
	// Convert the byte to a string
	var bString string = string(b)
	
	// Print the string representation
	fmt.Println("String :", bString)
	// Print the byte representation
	fmt.Println("Byte :", b)
}
