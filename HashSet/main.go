package main

import (
	"fmt"
	"hash/fnv"
)

type HashSet struct {
	set map[uint64]bool
}

// NewHashSet creates a new empty HashSet
func NewHashSet() *HashSet {
	return &HashSet{make(map[uint64]bool)}
}

// Add adds an element to the HashSet
func (hs *HashSet) Add(element []byte) {
	h := fnv.New64()
	h.Write(element)
	hs.set[h.Sum64()] = true
}

// Contains checks if an element is in the HashSet
func (hs *HashSet) Contains(element []byte) bool {
	h := fnv.New64()
	h.Write(element)
	_, ok := hs.set[h.Sum64()]
	return ok
}

func main() {
	// Create a new HashSet
	hs := NewHashSet()

	// Add some elements to the HashSet
	hs.Add([]byte("hello"))
	hs.Add([]byte("world"))
	hs.Add([]byte("golang"))

	// Check if the HashSet contains an element
	fmt.Println(hs.Contains([]byte("hello")))    // Output: true
	fmt.Println(hs.Contains([]byte("goodbye")))  // Output: false
}
