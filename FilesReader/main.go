package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// DONT FORGET TO CHANGE parentPath VALUE
var parentPath string = "C:/Users/Lenovo/OneDrive/Documents/Dev Go Project/FilesReader/"
var envPath    string = parentPath + ".env"
var txtPath    string = parentPath + "example.txt"
var jsonPath   string = parentPath + "example.json"

func main() {
	err := checkAllFiles()
	if err != nil {
		log.Fatal(err)
	}

	envData := readENV()
	fmt.Println(envData)

	jsonData := readJSON()
	fmt.Println(jsonData)

	txtData := readTXT()
	fmt.Println(txtData)
}

func checkAllFiles() error {
	_, err := os.Stat(envPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", envPath)
	}

	_, err = os.Stat(txtPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", txtPath)
	}

	_, err = os.Stat(jsonPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", jsonPath)
	}

	return nil
}

func readENV() map[string]string {
	// Open the file at the specified path
	file, err := os.Open(envPath)
	// If there is an error opening the file, log the error and exit
	if err != nil {
		log.Fatal(err)
	}
	// Defer closing the file until the function returns
	defer file.Close()

	// Initialize an empty map to store the key-value pairs
	lines := make(map[string]string)
	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)
	// Iterate over the lines in the file
	for scanner.Scan() {
		// Split the line by the "=" separator
		parts := strings.Split(scanner.Text(), "=")
		// If the line has exactly two parts (a key and a value)
		if len(parts) == 2 {
			// Add the key-value pair to the map
			lines[parts[0]] = parts[1]
		}
	}

	// If there is an error while scanning the file, log the error and exit
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the map of key-value pairs
	return lines
}

func readJSON() map[string]string {
	// Read the contents of the file at the specified path
	file, err := ioutil.ReadFile(jsonPath)
	// If there is an error reading the file, log the error and exit
	if err != nil {
		log.Fatal(err)
	}

	// Initialize an empty map to store the data
	var data map[string]string
	// Unmarshal the contents of the file into the map
	err = json.Unmarshal(file, &data)
	// If there is an error unmarshalling the data, log the error and exit
	if err != nil {
		log.Fatal(err)
	}

	// Return the map of data
	return data
}

func readTXT() []string {
	// Open the file at the specified path
	file, err := os.Open(txtPath)
	// If there is an error opening the file, log the error and exit
	if err != nil {
		log.Fatal(err)
	}
	// Defer closing the file until the function returns
	defer file.Close()

	// Initialize an empty slice of strings to store the lines
	var lines []string
	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)
	// Iterate over the lines in the file
	for scanner.Scan() {
		// Append each line to the slice of lines
		lines = append(lines, scanner.Text())
	}

	// If there is an error while scanning the file, log the error and exit
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the slice of lines
	return lines
}

