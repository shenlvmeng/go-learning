package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	if file == "" {
		fmt.Fprintln(os.Stderr, "err: missing filename")
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v\n", err)
		os.Exit(0)
	}
	wordsMap := countWords(f)
	fmt.Printf("string\t\tcount\n")
	for k, v := range wordsMap {
		fmt.Printf("%s\t\t%d\n", k, v)
	}
}

func countWords(f *os.File) map[string]int {
	wordsMap := make(map[string]int)
	in := bufio.NewScanner(f)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		word := in.Text()
		wordsMap[word]++
	}
	return wordsMap
}
