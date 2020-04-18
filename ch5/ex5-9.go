package main

import (
	"bytes"
	"fmt"
	"os"
)

const testString = "%foo"

func main() {
	testFunc := func(string) string {
		return "bar"
	}
	resStr := expand(os.Args[1], testFunc)
	fmt.Println(resStr)
}

func expand(s string, f func(string) string) string {
	var res bytes.Buffer
	b := []rune{}
	for _, r := range s {
		if byte(r) == testString[len(b)] {
			b = append(b, r)
			if string(b) == testString {
				resStr := f(testString)
				res.WriteString(resStr)
				b = []rune{}
			}
		} else {
			res.WriteString(string(b))
			res.WriteRune(r)
			b = []rune{}
		}
	}
	return res.String()
}
