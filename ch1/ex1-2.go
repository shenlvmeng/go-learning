package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", index, arg)
	}
}
