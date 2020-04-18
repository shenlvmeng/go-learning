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
	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		fmt.Println(n.FirstChild)
	}

	if c := n.FirstChild; c != nil {
		links = visit(links, c)
	} else if s := n.NextSibling; s != nil {
		links = visit(links, s)
	}

}
