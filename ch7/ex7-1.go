package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var w WordCounter
	var l LineCounter
	w, l = 0, 0

	fmt.Fprintln(&w, os.Args[1])
	fmt.Fprintln(&l, os.Args[1])
	fmt.Println(w)
	fmt.Println(l)
}

type WordCounter int

func (c *WordCounter) Write(b []byte) (int, error) {
	words, _, err := bufio.ScanWords(b, true)
	if err != nil {
		return 0, err
	}
	*c += WordCounter(words)
	return words, nil
}

type LineCounter int

func (c *LineCounter) Write(b []byte) (int, error) {
	lines, _, err := bufio.ScanLines(b, true)
	fmt.Print(lines, "line counter")
	if err != nil {
		return 0, err
	}
	*c += LineCounter(lines)
	return lines, nil
}
