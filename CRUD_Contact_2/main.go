package main

import (
	"errors"
	"fmt"
)

// Contact represents a single contact with a name and phone number.
type Contact struct {
	Name  string
	Phone string
}

// Contacts is a map of contacts, where the key is a unique ID for each contact.
type Contacts map[int]Contact

func main() {
	// vars
	var id int
	var contact Contact
	var err error

	// Create a new Contacts map.
	contacts := Contacts{}

	// Add a new contact to the Contacts map.
	id, err = contacts.AddContact("John Doe", "555-555-5555")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Contact added with ID", id)
	}

	// Get a contact from the Contacts map.
	contact, err = contacts.GetContact(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Contact name:", contact.Name)
		fmt.Println("Contact phone:", contact.Phone)
	}

	// Update the contact in the Contacts map.
	err = contacts.UpdateContact(id, "Jane Doe", "444-444-4444")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Contact updated")
	}

	err = contacts.DeleteContact(2) 
	// error -> contact[id] not found
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Contact Deleted")
	}

	err = contacts.DeleteContact(0) // success
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Contact Deleted")
	}
}

// AddContact adds a new contact to the Contacts map.
func (c *Contacts) AddContact(name, phone string) (int, error) {
	if name == "" || phone == "" {
		return 0, errors.New("> err : cannot add contact with empty name/ empty phone number")
	}

	// Generate a unique ID for the new contact.
	var id int
	for id = range *c {
		id++
	}

	// Add the new contact to the map.
	(*c)[id] = Contact{
		Name:  name,
		Phone: phone,
	}

	return id, nil
}

// GetContact retrieves a contact from the Contacts map by ID.
func (c *Contacts) GetContact(id int) (Contact, error) {
	if _, ok := (*c)[id]; !ok {
		return Contact{}, errors.New("> err : contact not found")
	}

	return (*c)[id], nil
}

// UpdateContact updates an existing contact in the Contacts map.
func (c *Contacts) UpdateContact(id int, name, phone string) error {
	if _, ok := (*c)[id]; !ok {
		return errors.New("> err : contact not found")
	}

	if name == "" || phone == "" {
		return errors.New("> err : cannot update contact with empty name/ empty phone number")
	}

	// Update the contact in the map.
	(*c)[id] = Contact{
		Name:  name,
		Phone: phone,
	}

	return nil
}

// DeleteContact deletes a contact from the Contacts map.
func (c *Contacts) DeleteContact(id int) error {
	if _, ok := (*c)[id]; !ok {
		return errors.New("> err : contact not found")
	}

	delete(*c, id)
	return nil
}