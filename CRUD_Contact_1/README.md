# Contact App 1

Implementing CRUD for contact-app in a `Contact struct` and methods for creating, reading, updating, and deleting contacts.

## Features

- Create new contacts
- Read existing contacts
- Update existing contacts
- Delete existing contacts

## Code structure

The code consists of the following main components:

The Contact struct, which has two fields: Name and Phone.
Methods that operate on Contact instances. These methods are defined using Go's receiver syntax, which allows them to be called using the `. operator` on Contact instances. For example, to create a new Contact, you can call the Create method like this:
