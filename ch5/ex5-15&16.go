package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(max(5, 4, 3, 2, 1))
	fmt.Println(min(5, 4, 3, 2, 1))
	fmt.Println(join("5", "4", "3", "2", "1"))
}

func max(vals ...float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	if len(vals) == 1 {
		return vals[0]
	}
	if len(vals) == 2 {
		return math.Max(vals[0]*1.0, vals[1]*1.0)
	}
	vals[1] = math.Max(vals[0], vals[1])
	return max(vals[1:]...)
}

func min(vals ...float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	if len(vals) == 1 {
		return vals[0]
	}
	if len(vals) == 2 {
		return math.Min(vals[0], vals[1])
	}
	vals[1] = math.Min(vals[0], vals[1])
	return min(vals[1:]...)
}

func join(strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 2 {
		return strs[0] + strs[1]
	}
	strs[1] = strs[0] + strs[1]
	return join(strs[1:]...)
}
