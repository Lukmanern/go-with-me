package main

import "fmt"

func main() {
	compare()
	byte_to_string()
}

func compare() {
	fmt.Println(`string => "a" :`, "a")
	fmt.Println(`byte => 'a' :`, 'a')
}

func byte_to_string() {
	var b byte = 'b'
	var bString string = string(b)
	
	fmt.Println("String :", bString, "\nByte :", b)
}