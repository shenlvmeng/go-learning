package main

import (
	"fmt"
	"os"
)

func deDup(strs []string) []string {
	res := []string{}
	var last string
	for index, n := range strs {
		if n == last && index != 0 {
			continue
		}
		res = append(res, n)
		last = n
	}
	return res
}

func main() {
	fmt.Printf("%v\n", deDup(os.Args[1:]))
}
