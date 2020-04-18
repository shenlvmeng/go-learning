package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	"linear algebra":        {"calculus"},
}

func main() {
	res, cycles := topoSort(prereqs)
	for i, course := range res {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Println("Cycles: ", cycles)
}

func topoSort(m map[string][]string) ([]string, int) {
	var order []string
	seen := make(map[string]bool)
	cycles := 0

	var visitAll func(items []string, target string)
	visitAll = func(items []string, target string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if target == "" {
					visitAll(m[item], item)
				} else {
					visitAll(m[item], target)
				}
				order = append(order, item)
			} else {
				cycles++
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys, "")
	return order, cycles
}
