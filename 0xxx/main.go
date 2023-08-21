package main

import "fmt"

// SEC_COMMIT: This constant has a hexadecimal value 0x08000000.
// It is likely used to indicate that a memory section should be committed,
// which means that the memory associated with the section should be allocated for use.

// SECTION_WRITE: This constant has a value 0x2 (which is 2 in decimal).
// It's used to indicate that the memory section is writable, meaning data can be written to it.

// SECTION_READ: This constant has a value 0x4 (which is 4 in decimal).
// It's used to indicate that the memory section is readable, meaning data can be read from it.

// SECTION_EXECUTE: This constant has a value 0x8 (which is 8 in decimal).
// It's used to indicate that the memory section is executable, meaning code can be executed from it.

// SECTION_RWX: This constant is a combination of the SECTION_WRITE,
// SECTION_READ, and SECTION_EXECUTE constants. It's created by using
// bitwise OR (|) to combine the individual constants, resulting in a
// value of 0xE (which is 14 in decimal). This can be used to indicate
// a memory section that is writable, readable, and executable.

const (
	SEC_COMMIT      = 0x08000000
	SECTION_WRITE   = 0x2
	SECTION_READ    = 0x4
	SECTION_EXECUTE = 0x8
	SECTION_RWX     = SECTION_WRITE | SECTION_READ | SECTION_EXECUTE
)

func main() {
	fmt.Printf("SEC_COMMIT: %d\n", SEC_COMMIT)
	fmt.Printf("SECTION_WRITE: %d\n", SECTION_WRITE)
	fmt.Printf("SECTION_READ: %d\n", SECTION_READ)
	fmt.Printf("SECTION_EXECUTE: %d\n", SECTION_EXECUTE)
	fmt.Printf("SECTION_RWX: %d\n", SECTION_RWX)
}
