package main

import (
	"fmt"
	"unicode"
)

func main() {
	// The message to be encrypted
	message := "HELLO, WORLD!"

	// The key to use for the cipher
	key := 3

	// Encrypt the message
	encrypted := encryptCaesar(message, key)
	fmt.Println("Encrypted message:", encrypted)

	// Decrypt the message
	decrypted := decryptCaesar(encrypted, key)
	fmt.Println("Decrypted message:", decrypted)
}

func encryptCaesar(plaintext string, key int) string {
	// Convert the plaintext to a slice of 
	// runes (a slice of Unicode code points)
	runes := []rune(plaintext)

	// Shift each letter in the 
	// plaintext by the key
	for i, r := range runes {
		// Shift only letters, not digits 
		// or other characters
		if unicode.IsLetter(r) {
			// Shift the letter
			runes[i] = shiftRune(r, key)
		}
	}

	// Convert the slice of runes 
	// back to a string and return it
	return string(runes)
}

func decryptCaesar(ciphertext string, key int) string {
	// Convert the ciphertext to a slice of 
	// runes (a slice of Unicode code points)
	runes := []rune(ciphertext)

	// Shift each letter in the 
	// ciphertext by the key
	for i, r := range runes {
		// Shift only letters, not digits 
		// or other characters
		if unicode.IsLetter(r) {
			// Shift the letter
			runes[i] = shiftRune(r, -key)
		}
	}

	// Convert the slice of runes back 
	// to a string and return it
	return string(runes)
}

func shiftRune(r rune, key int) rune {
	// Convert the rune to an integer
	s := int(r)

	// Shift the letter by the key, keeping 
	// it within the bounds of the alphabet
	if unicode.IsLower(r) {
		// If the letter is lowercase, 
		// shift it within the
		// range 'a' to 'z'
		s = (s - 'a' + key) % 26 + 'a'
	} else if unicode.IsUpper(r) {
		// If the letter is uppercase, 
		// shift it within the range 'A' to 'Z'
		s = (s - 'A' + key) % 26 + 'A'
	}

	// Convert the shifted integer 
	// back to a rune and return it
	return rune(s)
}

