package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	input, flag := os.Args[1], os.Args[2]
	var sum []uint8
	if flag == "384" {
		b := sha512.Sum384([]byte(input))
		sum = b[:]
	} else if flag == "512" {
		b := sha512.Sum512([]byte(input))
		sum = b[:]
	} else {
		b := sha256.Sum256([]byte(input))
		sum = b[:]
	}
	fmt.Printf("%x\n", sum)
}
