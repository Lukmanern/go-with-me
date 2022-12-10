package main

import (
	"crypto/sha256"
	"fmt"
)

func encryptPassword(password string) string {
	// Hash the password using SHA-256
	h := sha256.New()
	h.Write([]byte(password))
	encryptedPassword := h.Sum(nil)

	// Return the hexadecimal representation of the hash
	return fmt.Sprintf("%x", encryptedPassword)
}

func main() {
	// Encrypt the password "secret_password"
	encryptedPassword := encryptPassword("secret_password")
	fmt.Println(encryptedPassword)
}
