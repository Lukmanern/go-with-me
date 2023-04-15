package main

import (
	"fmt"
)

type Contact struct {
	Name  string
	Phone string
}

func main() {
	var c Contact
	c.Create("John Doe", "555-555-5555")
	fmt.Println("Created contact:")
	fmt.Println(c.Read())

	c.Update("Jane Doe", "444-444-4444")
	fmt.Println("Updated contact:")
	fmt.Println(c.Read())

	c.Delete()
	fmt.Println("Deleted contact:")
	fmt.Println(c.Read())
}

// Create a new contact
func (c *Contact) Create(name, phone string) {
	c.Name = name
	c.Phone = phone
}

// Read the contact's name and phone number
func (c *Contact) Read() (string, string) {
	return c.Name, c.Phone
}

// Update the contact's name and phone number
func (c *Contact) Update(name, phone string) {
	c.Name = name
	c.Phone = phone
}

// Delete the contact (null-ing values)
func (c *Contact) Delete() {
	c.Name = ""
	c.Phone = ""
}
