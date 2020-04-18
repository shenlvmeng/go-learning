package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	var allDupFiles []string
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup err: %v\n", err)
			continue
		}

		if status := countlines(f, counts); status == 1 {
			allDupFiles = append(allDupFiles, file)
		} else if status == 2 {
			allDupFiles = []string{file}
		} else if status == 3 {
			allDupFiles = []string{}
		}
		f.Close()
	}
	fmt.Println(strings.Join(allDupFiles, "\n"))
}

// 0: 不变 1: 添加 2: 取代 3: 清零
func countlines(f *os.File, counts map[string]int) int {
	input := bufio.NewScanner(f)
	status := 1
	originalLen := len(counts)
	thisCounts := make(map[string]int)
	for input.Scan() {
		text := input.Text()
		_, ok := counts[text]
		if !ok {
			status = 2
		} else {
			thisCounts[text]++
		}
		counts[text]++
	}
	if status == 2 && len(thisCounts) < originalLen {
		status = 3
	} else if len(thisCounts) < originalLen {
		status = 0
	}
	return status
}
