package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	printTextNodes(doc)
}

// printTextNodes prints the contents of all text nodes in an HTML document tree.
func printTextNodes(n *html.Node) {
	if n == nil || n.Data == "script" || n.Data == "style" {
		return
	}
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	printTextNodes(n.FirstChild)
	printTextNodes(n.NextSibling)
}
