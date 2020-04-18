package main

import "fmt"

func rotate(s []int, pivot int) []int {
	res := make([]int, len(s))
	copy(res, s[pivot:])
	copy(res[len(s)-pivot:], s[:pivot])
	return res
}

func main() {
	arr := [...]int{1, 2, 3, 5, 6, 2, 4}
	fmt.Printf("%v\n", rotate(arr[:], 2))
}
