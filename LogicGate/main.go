package main

import "fmt"

// LogicGate represents a basic logic gate.
type LogicGate func(inputs ...bool) bool

// NOT gate implementation.
func NOT(input bool) bool {
	return !input
}

// AND gate implementation.
func AND(inputs ...bool) bool {
	result := true
	for _, input := range inputs {
		result = result && input
	}
	return result
}

// OR gate implementation.
func OR(inputs ...bool) bool {
	result := false
	for _, input := range inputs {
		result = result || input
	}
	return result
}

// XOR gate implementation.
func XOR(inputs ...bool) bool {
	count := 0
	for _, input := range inputs {
		if input {
			count++
		}
	}
	return count%2 == 1
}

// XNOR gate implementation.
func XNOR(inputs ...bool) bool {
	return !XOR(inputs...)
}

// NOR gate implementation.
func NOR(inputs ...bool) bool {
	return !OR(inputs...)
}

func main() {
	inputA := true
	inputB := false

	fmt.Printf("NOT(%v) = %v\n", inputA, NOT(inputA))
	fmt.Printf("AND(%v, %v) = %v\n", inputA, inputB, AND(inputA, inputB))
	fmt.Printf("OR(%v, %v) = %v\n", inputA, inputB, OR(inputA, inputB))
	fmt.Printf("XOR(%v, %v) = %v\n", inputA, inputB, XOR(inputA, inputB))
	fmt.Printf("XNOR(%v, %v) = %v\n", inputA, inputB, XNOR(inputA, inputB))
	fmt.Printf("NOR(%v, %v) = %v\n", inputA, inputB, NOR(inputA, inputB))
}
