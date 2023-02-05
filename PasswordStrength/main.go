package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(checkPasswordStrength("P@ssword1"))	// Output: strong
	fmt.Println(checkPasswordStrength("password"))	// Output: weak
	fmt.Println(checkPasswordStrength("Passw0rd"))	// Output: medium
	fmt.Println(checkPasswordStrength("p@ssword"))	// Output: medium
}

func checkPasswordStrength(password string) string {
	var strength string
	var numUpper, numLower, numDigits, numSpecial int

	// check for length of password
	if len(password) < 8 {
		strength = "weak"
		return strength
	}

	// check for the presence of uppercase letters
	uppercase := regexp.MustCompile(`[A-Z]`)
	numUpper = len(uppercase.FindAllString(password, -1))

	// check for the presence of lowercase letters
	lowercase := regexp.MustCompile(`[a-z]`)
	numLower = len(lowercase.FindAllString(password, -1))

	// check for the presence of digits
	digits := regexp.MustCompile(`[0-9]`)
	numDigits = len(digits.FindAllString(password, -1))

	// check for the presence of special characters
	special := regexp.MustCompile(`[!@#\$%\^&\*\(\)]`)
	numSpecial = len(special.FindAllString(password, -1))

	// evaluate password strength
	if numUpper > 0 && numLower > 0 && numDigits > 0 && numSpecial > 0 {
		strength = "strong"
	} else if (numUpper > 0 && numLower > 0 && numDigits > 0) || (numUpper > 0 && numLower > 0 && numSpecial > 0) || (numUpper > 0 && numSpecial > 0 && numDigits > 0) || (numLower > 0 && numSpecial > 0 && numDigits > 0) {
		strength = "medium"
	} else {
		strength = "weak"
	}
	
	return strength
}