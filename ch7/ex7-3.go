package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	res := t.left.String() + fmt.Sprintf("%d", t.value) + t.right.String()
	return res
}

func buildTree(data []int) *tree {
	root := new(tree)
	for _, v := range data {
		root = add(root, v)
	}
	return root
}

func add(root *tree, value int) *tree {
	if root == nil {
		root = new(tree)
		root.value = value
		return root
	}

	if value < root.value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	root := buildTree(data)
	fmt.Println(root)

	fmt.Println(new(tree))
}
