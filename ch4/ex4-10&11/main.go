package main

import (
	"os"
)

func main() {
	res, err := searchIssues(os.Args[1:])
}
