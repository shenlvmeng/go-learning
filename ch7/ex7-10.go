package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 2, 1}

	fmt.Println(IsPalindrome(sort.IntSlice(a)))
	fmt.Println(IsPalindrome(sort.IntSlice(b)))
}

func IsPalindrome(s sort.Interface) bool {
	if s.Len() == 0 {
		return false
	}
	for i, j := 0, s.Len()-1; i < j; {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
