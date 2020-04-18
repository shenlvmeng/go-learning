package main

import "io"

func main() {

}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CountWriter{
		rawWriter: w,
	}
	count := cw.Count
	return &cw, &count
}

type CountWriter struct {
	rawWriter io.Writer
	Count     int64
}

func (c *CountWriter) Write(b []byte) (int, error) {
	byteCount, err := c.rawWriter.Write(b)
	c.Count += int64(len(b))

	return byteCount, err
}
