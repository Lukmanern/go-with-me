package main

import (
	"fmt"
)

// In Go, `iota` is a predeclared identifier that
// represents successive untyped integer constants.
// It is used to create a sequence of related constants
// with a common prefix.
// iota is a useful tool for creating sequences of related
// constants in Go. It allows you to avoid having to specify
// the values of the constants manually, and it makes it easy
// to add or remove constants from the sequence.
const (
	VERY_LOW = iota
	LOW
	MEDIUM
	HIGH
	VERY_HIGH
)

type FuzzySet struct {
	label string
	value int
}

func (fs FuzzySet) String() string {
    	return fmt.Sprintf("%s: %d", fs.label, fs.value)
}

func Fuzzify(value int, sets []FuzzySet) []FuzzySet {
	for i := range sets {
		switch {
		case value < 25:
			sets[i].value = VERY_LOW
		case value < 50:
			sets[i].value = LOW
		case value < 75:
			sets[i].value = MEDIUM
		case value < 100:
			sets[i].value = HIGH
		default:
			sets[i].value = VERY_HIGH
		}
	}

	return sets
}

func main() {
	sets := []FuzzySet{
		{"Temperature", 0},
		{"Pressure", 0},
		{"Humidity", 0},
	}

	// Fuzzify the values
	sets = Fuzzify(75, sets)

	// Print the fuzzy sets
	for _, set := range sets {
		fmt.Println(set)
	}
}
