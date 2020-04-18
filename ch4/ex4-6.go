package main

import (
	"fmt"
	"os"
	"unicode"
)

func squash(b []byte) []byte {
	s := string(b)
	bytes := []byte{}
	for _, char := range s {
		if !unicode.IsSpace(char) {
			bytes = append(bytes, byte(char))
		}
	}
	return bytes
}

func main() {
	str := os.Args[1]
	fmt.Printf("%s\n", squash([]byte(str)))
}
