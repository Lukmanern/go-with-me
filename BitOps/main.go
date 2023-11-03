package main

import "fmt"

var (
	data1, data2 uint16
)

func print(i uint16) {
	fmt.Printf("%08b\n", i)
}

func main() {
	and()
	or()
	shift()
	xor()
	bitwise()
	nor()
}

// AND &
func and() {
	data1 = 0b_0101
	data2 = 0b_1010

	print(data1 & data2) // => 00000000
}

func or() {
	data1 = 0b_0101
	data2 = 0b_1010

	print(data1 | data2) // => 00001111
}

func shift() {
	data1 = 0b_0101
	data2 = 0b_1010

	print(data1 << data2) // => 1010000000000
	print(data1 >> data2) // => 00000000
}

func xor() {
	data1 = 0b_0101
	data2 = 0b_1010

	print(data1 ^ data2) // => 00001111
}

func bitwise() {
	data1 = 0b_0101
	data2 = 0b_1010

	print(data1 &^ data2) // => 00000101 == data1
}

func nor() {
	data1 = 0b_0101
	data2 = 0b_1010
	result := ^(data1 | data2)
	print(result) // => 1111111111110000
}
