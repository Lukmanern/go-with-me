# Regular Expression Matching

This is a simple Go program that demonstrates how to use the `regexp` package to match a regular expression pattern in a string.

The program defines a main function that creates a string `pattern` and a string text. The main function then uses the `regexp.Compile` function to create a regular expression object from the pattern string. The program then uses the MatchString method of the regular expression object to search for the pattern in the text.

In this example, the program searches for the string "foo" in the text "My text contains a pattern: foo". If the pattern is found, the program prints "Pattern found in text", otherwise it prints "Pattern not found in text".
