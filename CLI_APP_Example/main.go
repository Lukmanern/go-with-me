package main

import (
	"fmt"
	"os"
)

func main() {
	// Define a map of supported commands 
	// and their corresponding functions
	commands := map[string]func([]string){
		"hello": hello,
		"bye":   bye,
	}

	// Get the command line 
	// arguments passed to the program
	args := os.Args[1:]

	// Check if no command was specified
	if len(args) == 0 {
		fmt.Println("Error: no command specified")
		return
	}

	// Get the command and the 
	// remaining arguments
	command := args[0]
	args = args[1:]

	// Check if the command is supported (or exist)
	if fn, ok := commands[command]; ok {
		// Execute the corresponding function
		fn(args)
	} else {
		fmt.Printf("Error: unsupported command '%s'\n", command)
	}
}

// Define functions for each 
// supported command

func hello(args []string) {
	// Print a simple message
	fmt.Println("Hello, world!")

	// Print the remaining arguments, if any
	if len(args) > 0 {
		fmt.Println("Arguments:", args)
	}
}

func bye(args []string) {
	// Print a simple message
	fmt.Println("Goodbye, world!")

	// Print the remaining arguments, if any
	if len(args) > 0 {
		fmt.Println("Arguments:", args)
	}
}
