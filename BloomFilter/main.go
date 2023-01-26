package main

import (
	"fmt"
	"hash"
	"hash/fnv"
)

type BloomFilter struct {
	bits []bool
	k    uint
	h    hash.Hash64
}

// NewBloomFilter creates a new Bloom filter with 
// the given number of bits and hash functions
func NewBloomFilter(m uint, k uint) *BloomFilter {
	bits := make([]bool, m)
	h := fnv.New64()
	return &BloomFilter{bits, k, h}
}

// Add adds an element to the Bloom filter
func (bf *BloomFilter) Add(element []byte) {
	for i := uint(0); i < bf.k; i++ {
		bf.h.Reset()
		bf.h.Write(element)
		bf.h.Write([]byte(fmt.Sprintf("%d", i)))
		index := bf.h.Sum64() % uint64(len(bf.bits))
		bf.bits[index] = true
	}
}

// Contains checks if an element 
// is possibly in the Bloom filter
func (bf *BloomFilter) Contains(element []byte) bool {
	for i := uint(0); i < bf.k; i++ {
		bf.h.Reset()
		bf.h.Write(element)
		bf.h.Write([]byte(fmt.Sprintf("%d", i)))
		index := bf.h.Sum64() % uint64(len(bf.bits))
		if !bf.bits[index] {
			return false
		}
	}
	return true
}

func main() {
	// Create a new Bloom filter with 100 bits and 5 hash functions
	bf := NewBloomFilter(100, 5)

	// Add some elements to the Bloom filter
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("bloom filter"))

	// Check if the Bloom filter contains an element
	fmt.Println(bf.Contains([]byte("hello")))    // Output: true
	fmt.Println(bf.Contains([]byte("goodbye")))  // Output: false
}
