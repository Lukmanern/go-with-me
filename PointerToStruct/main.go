package main

import "fmt"

// Define a new type called "Person"
type Person struct {
	// Fields within the struct
	name string
	age  int
	// The address field
	// is a pointer to another struct
	address *Address
}

// Define a new type called "Address"
type Address struct {
	// Fields within the struct
	street string
	city   string
	state  string
	zip    string
}

func main() {
	// Create a new instance
	// of the Address type
	a := Address{
		street: "123 Main St",
		city:   "Anytown",
		state:  "CA",
		zip:    "98765",
	}

	// Create a new instance
	//  of the Person type
	p := Person{
		name:    "John Doe",
		age:     29,
		address: &a,
	}

	// Becouse of address is a
	// pointer, we can change address
	// of person by change
	// value of `a` variable
	a.state = "NY"

	// Print the values of the
	// fields within the struct
	fmt.Printf("%s is %d years old and lives at %s, %s, %s %s\n",
		p.name, p.age, p.address.street,
		p.address.city, p.address.state,
		p.address.zip,
	)
}
