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
	if n == nil {
		return
	}
	if n.Data == "script" || n.Data == "style" { // do not descend into <script> or <style> elements, since their contents are not visible in a web browser
		return
	}
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	printTextNodes(n.FirstChild)
	printTextNodes(n.NextSibling)
}
