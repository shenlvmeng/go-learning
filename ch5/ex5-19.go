package main

import "fmt"

func main() {
	fmt.Print(p())
}

func p() (val int) {
	defer func() {
		p := recover()
		if p != nil {
			val = 1
		}
	}()

	panic("some what")
}
