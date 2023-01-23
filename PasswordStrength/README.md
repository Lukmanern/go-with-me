# Password Strength Checker

This is a simple Go program that checks the strength of a given password using regular expressions.

The program defines a `checkPasswordStrength` function which takes in a string (password) as input. The function uses regular expressions to check for the presence of different types of characters in the password and assigns a strength value based on the presence of those characters.

The function checks for the presence of uppercase letters, lowercase letters, digits and special characters. If a password contains all of these types of characters, the password is considered `strong`. If a password contains at least three types of characters, the password is considered `medium`, otherwise, the password is considered `weak`.
