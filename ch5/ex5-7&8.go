package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	printHTML(os.Args[1], testGetElementByID)
}

func printHTML(url string, test func(doc *html.Node)) {
	root, err := fetchHTML(url)
	if err != nil {
		log.Fatalf("fetching %s %v", url, err)
	}
	printDoc(root)
	test(root)
}

func testGetElementByID(doc *html.Node) {
	fmt.Println("Now input element you want to search...\n")
	// get input
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	id := in.Text()
	if id == "" {
		fmt.Println("Empty input. Bye.")
		return
	}
	node := elementByID(doc, id)
	if node != nil {
		fmt.Printf("Getcha! A %s!\n", node.Data)
		return
	}
}

func fetchHTML(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing html: %s", err)
	}
	return doc, nil
}

func printDoc(root *html.Node) {
	depth := 0
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			for _, attr := range n.Attr {
				fmt.Printf(" %s=\"%s\"", attr.Key, attr.Val)
			}
			fmt.Println(">\n")
		}
		depth++
		return false
	}
	post := func(n *html.Node) {
		depth--
		if n.Type == html.ElementNode {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	handleSingle := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s />\n", depth*2, "", n.Data)
		} else if n.Type == html.TextNode {
			// fmt.Printf("%*s%s", depth*2, "", n.Data)
		}
	}

	traverse(root, pre, post, handleSingle)
}

func elementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == id {
					return true
				}
			}
		}
		return false
	}
	return traverse(doc, pre, nil, nil)
}

func traverse(node *html.Node, in func(node *html.Node) bool, out func(node *html.Node), onSingle func(node *html.Node)) (res *html.Node) {
	if node.FirstChild == nil {
		if onSingle != nil {
			onSingle(node)
		}
		res = nil
	}
	if in != nil && in(node) {
		res = node
		return
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		res = traverse(c, in, out, onSingle)
		if res != nil {
			break
		}
	}

	if out != nil {
		out(node)
	}

	return
}
