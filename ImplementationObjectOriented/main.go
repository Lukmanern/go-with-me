package main

import "fmt"

// Define a struct that represents a person.
// Person define as Object
type Person struct { 
    // The person's name.
    Name string
    // The person's age.
    Age int
}

// Define a method for the Person struct that 
// allows a person to introduce themselves.
// Pointer of Person make Introduce() define as METHOD
func (p *Person) Introduce() { 
    fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // Create a new instance of the Person struct.
    p := Person{Name: "John", Age: 30}

    // Call the Introduce method on the person instance.
    p.Introduce()
}