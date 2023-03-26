package main

import (
	"fmt"
	"unsafe"
)

// In this case, use the unsafe.Pointer to obtain a pointer
// to the first element of the x slice, and then use pointer
// arithmetic to iterate over the elements of the slice.
// Also use the *int type conversion to dereference
// the pointer and obtain the value of each element.

func main() {
	var x = []int{1, 2, 3, 4}

	// get pointer to the first element in the slice
	p := unsafe.Pointer(&x[0])

	for i := 0; i < len(x); i++ {
		// create a pointer to the i-th element in the slice
		ptr := unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(x[0]))

		// convert the pointer to a pointer to an int
		val := *(*int)(ptr)

		fmt.Printf("x[%d] = %d\n", i, val)
	}
}
