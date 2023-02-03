package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
)

func main() {
	// Store the password hash
	hashedPassword := hashPassword("secret_password")
	fmt.Println("hashed :", hashedPassword)
	// Validate the entered password
	if validatePassword("secret_password", hashedPassword) {
		fmt.Println("Password is valid.")
	} else {
		fmt.Println("Password is invalid.")
	}
}

func hashPassword(password string) []byte {
	// Hash the password using SHA-256
	h := sha256.New()
	h.Write([]byte(password))
	return h.Sum(nil)
}

func validatePassword(password string, hashedPassword []byte) bool {
	// Hash the entered password
	enteredHash := hashPassword(password)

	// Compare the two hashes
	return subtle.ConstantTimeCompare(enteredHash, hashedPassword) == 1
}
