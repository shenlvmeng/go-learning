package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func alg1(args []string) {
	s, sep := "", ""
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func alg2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
}

func alg3(args []string) {
	var _ = strings.Join(args[1:], " ")
}

func main() {
	start := time.Now()
	for i := 1; i < 1000; i++ {
		alg1(os.Args)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("arg1: %.6fs\n", secs)

	start = time.Now()
	for i := 1; i < 1000; i++ {
		alg2(os.Args)
	}
	secs = time.Since(start).Seconds()
	fmt.Printf("arg1: %.6fs\n", secs)

	start = time.Now()
	for i := 1; i < 1000; i++ {
		alg3(os.Args)
	}
	secs = time.Since(start).Seconds()
	fmt.Printf("arg1: %.6fs\n", secs)
}
