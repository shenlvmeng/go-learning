package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := countWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatalf("count failed: %v\n", err)
	}
	fmt.Printf("Words\tImages\n")
	fmt.Printf("%d\t%d\n", words, images)
}

func countWordsAndImages(url string) (words, images int, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s\n", err)
		return
	}
	words, images = count(doc)
	return
}

func count(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		if n.Data == "script" || n.Data == "style" {
			words = 0
			images = 0
			return
		}
		if n.Data == "img" {
			words = 0
			images = 1
			return
		}
	} else if n.Type == html.TextNode {
		s := strings.Fields(n.Data)
		words = len(s)
		images = 0
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		subWords, subImages := count(c)
		words += subWords
		images += subImages
	}
	return
}
