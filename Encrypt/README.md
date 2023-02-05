### Password Encryption

This is a simple Go program that demonstrates how to encrypt a password using the SHA-256 cryptographic hash function.

### Implementation Details

The program uses the `crypto/sha256` package from the Go standard library to compute the SHA-256 hash of a password. The encryptPassword function takes a password as input, hashes it using the `sha256.New` function, and then returns the hexadecimal representation of the resulting hash. The main function demonstrates how to use the `encryptPassword` function to encrypt the password "secret_password".

### Security

SHA-256 is a widely used cryptographic hash function that is believed to be secure for most purposes. However, it is important to note that this program is only intended to demonstrate the basics of password encryption in Go and should not be used in production systems without further modifications to improve security. For example, a production-ready password encryption system should use a strong and unique salt for each password, and should use a key derivation function such as [PBKDF2](https://en.wikipedia.org/wiki/PBKDF2) to increase the computational cost of attempting to brute-force the encrypted passwords.
