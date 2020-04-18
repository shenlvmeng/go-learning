package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func diffInt(i1 byte, i2 byte) int {
	var diffBits int
	for index := range [...]int{7: 1} {
		var diffUint int
		if ((i1 & (1 << uint(index))) & (i2 & (1 << uint(index)))) > 0 {
			diffUint = 1
		}
		diffBits += diffUint
	}
	return diffBits
}

func diffBits(b1 [32]byte, b2 [32]byte) int {
	var diffBits int
	for index := range b1 {
		diffBits += diffInt(b1[index], b2[index])
	}
	return diffBits
}

func main() {
	c1, c2 := sha256.Sum256([]byte(os.Args[1])), sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("Different bits: %d\n", diffBits(c1, c2))
}
