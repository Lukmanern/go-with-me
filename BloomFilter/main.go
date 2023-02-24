package main

import (
	"fmt"
	"hash"
	"hash/fnv"
)

// BloomFilter is a data structure that is
// used to test whether an element is a member of a set.
// It is probabilistic, meaning that it may produce false
// positives, but never false negatives.
type BloomFilter struct {
	bits []bool // slice of bits used in the Bloom filter
	k    uint   // number of hash functions to use
	h    hash.Hash64 // hash function
}

// NewBloomFilter creates a new Bloom filter 
// with the given number of bits and hash functions
func NewBloomFilter(m uint, k uint) *BloomFilter {
	bits := make([]bool, m)
	h := fnv.New64()
	return &BloomFilter{bits, k, h}
}

// Add adds an element to the Bloom filter
func (bf *BloomFilter) Add(element []byte) {
	for i := uint(0); i < bf.k; i++ {
		// reset hash
		bf.h.Reset() 
		// write element to hash
		bf.h.Write(element) 
		// write i to hash
		bf.h.Write([]byte(fmt.Sprintf("%d", i))) 
		// get index by taking the mod of the sum of the hash
		index := bf.h.Sum64() % uint64(len(bf.bits)) 
		// set the bit at the index to true
		bf.bits[index] = true 
	}
}

// Contains checks if an element is 
// possibly in the Bloom filter
func (bf *BloomFilter) Contains(element []byte) bool {
	for i := uint(0); i < bf.k; i++ {
		// reset hash
		bf.h.Reset() 
		// write element to hash
		bf.h.Write(element) 
		// write i to hash
		bf.h.Write([]byte(fmt.Sprintf("%d", i))) 
		// get index by taking the mod of the sum of the hash
		index := bf.h.Sum64() % uint64(len(bf.bits)) 
		// if the bit at the index is not set
		if !bf.bits[index] { 
			// the element is not in the filter
			return false 
		}
	}
	// all bits are set, element 
	// is possibly in the filter
	return true 
}


func main() {
	// Create a new Bloom filter with 
	// 100 bits and 5 hash functions
	bf := NewBloomFilter(100, 5)

	// Add some elements to the Bloom filter
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("bloom filter"))

	// Check if the Bloom filter contains an element
	fmt.Println(bf.Contains([]byte("hello")))    // Output: true
	fmt.Println(bf.Contains([]byte("goodbye")))  // Output: false
}
