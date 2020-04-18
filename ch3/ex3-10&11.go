package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	b := []byte(s)
	dotPos, end := strings.Index(s, "."), len(s)

	if dotPos > 0 {
		end = dotPos
	}
	var buffer bytes.Buffer
	for index, bt := range b {
		buffer.WriteByte(bt)
		if index < end-1 && (end-index-1)%3 == 0 {
			buffer.WriteRune(',')
		}
	}
	return buffer.String()
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
