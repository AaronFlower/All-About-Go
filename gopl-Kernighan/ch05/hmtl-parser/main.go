package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://aaronflower.github.io/about/"
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	// scanner := bufio.NewScanner(resp.Body)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	traverseNode(doc, 0)
}

func traverseNode(node *html.Node, level int) {
	pad := strings.Repeat(" ", level*4)
	if node.Type == html.ElementNode {
		fmt.Printf("%s<%s> \n", pad, node.Data)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		traverseNode(c, level+1)
	}
	if node.Type == html.ElementNode {
		fmt.Printf("%s</%s> \n", pad, node.Data)
	}
}
