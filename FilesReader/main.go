package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// SABAR BELOM JADI WKWKWKW
func main() {
	filesPath := map[string]string{
		"txt" : "DONT_FORGET_TO_CHANGE_THIS/Files_Reader/text.txt",
	}
	for key, value := range filesPath {
		if isFileExist(filesPath[key]) {
			readTXT(value)
		}
	}
}

func isFileExist(filePath string) bool {
	// Path to the file we want to check
	// Use os.Stat to check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// The file doesn't exist
		fmt.Println(err)
		return false
	}
	return true
}

func readTXT(filePath string) {
	// Path to the file we want to read
	// Read the file
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		// There was an error reading the file
		fmt.Println(err)
		return
	}
	// Convert the bytes to a string
	text := string(bytes)
	// Print the file contents
	fmt.Println(text)
}