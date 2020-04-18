package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLink: %v\n", err)
		os.Exit(1)
	}
	nodeMap := make(map[string]int)
	fmt.Printf("type\t\tcount\n")
	for nodeName, count := range visit(nodeMap, doc) {
		fmt.Printf("%s\t\t%d\n", nodeName, count)
	}
}

func visit(nodeMap map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		nodeMap[n.Data]++
	}
}
