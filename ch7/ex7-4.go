package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	data  string
	index int
}

func (sr *StringReader) Read(p []byte) (int, error) {
	data := []byte(sr.data)
	if sr.index >= len(data) {
		return 0, io.EOF
	}

	data = data[sr.index:]
	n := 0

	n = copy(p, data)
	sr.index += n
	return n, nil
}

func NewReader(in string) *StringReader {
	sr := StringReader{
		data: in,
	}
	return &sr
}

func main() {
	str := "Hello world!"
	sr := NewReader(str)
	data := make([]byte, 10)
	n, err := sr.Read(data)
	for err == nil {
		fmt.Println(n, string(data[0:n]))
		n, err = sr.Read(data)
	}
}
