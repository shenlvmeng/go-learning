package main

import (
	"fmt"
	"io"
	"os"
)

type LimitedReader struct {
	Reader  io.Reader
	Limit   int
	current int
}

func (lr *LimitedReader) Read(b []byte) (int, error) {
	if lr.current >= lr.Limit {
		return 0, io.EOF
	}

	if lr.current+len(b) > lr.Limit {
		b = b[:lr.Limit-lr.current]
	}
	n, err := lr.Reader.Read(b)
	if err != nil {
		return n, err
	}
	lr.current += n
	return n, nil
}

func LimitReader(r io.Reader, n int) io.Reader {
	lr := LimitedReader{
		Reader: r,
		Limit:  n,
	}
	return &lr
}

func main() {
	file, err := os.Open("7-5.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	lr := LimitReader(file, 6)
	buf := make([]byte, 10)
	n, err := lr.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, buf)
}
