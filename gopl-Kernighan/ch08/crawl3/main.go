package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AaronFlower/All-About-Go/gopl-Kernighan/ch05/links"
)

func crawl(link string) []string {
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenlinks := make(chan string) // de-duplicated URLs

	// Start with the command-line arguments.
	go func() {
		worklist <- os.Args[1:]
	}()

	// create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }() // 防止被 blocked
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- link
			}
		}
	}
}
