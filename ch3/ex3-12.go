package main

import (
	"fmt"
	"os"
)

func testAnagram(s1 string, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)

	if len(s1) != len(s2) {
		return false
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for index, r := range b1 {
		m1[r]++
		m2[b2[index]]++
	}

	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "anagram: need 2 strings")
		os.Exit(1)
	}
	fmt.Printf("%v", testAnagram(os.Args[1], os.Args[2]))
}
