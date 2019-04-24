package main

import (
	"flag"
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

/**
 * usage: go run main.go -depth=3 http://gopl.io/
 */
func main() {
	depthPtr := flag.Int("depth", 2, "The crawling depth")
	flag.Parse()
	type depthlink struct {
		depth int
		links []string
	}
	worklist := make(chan depthlink)    // lists of URLs, may have duplicates
	unseenlinks := make(chan depthlink) // de-duplicated URLs

	// Start with the command-line arguments.
	go func() {
		l := len(os.Args)
		worklist <- depthlink{0, os.Args[l-1:]}
	}()

	// create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlinks {
				foundLinks := crawl(link.links[0])
				go func(link depthlink) { worklist <- depthlink{link.depth + 1, foundLinks} }(link) // 防止被 blocked
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		if list.depth <= *depthPtr {
			for _, link := range list.links {
				if !seen[link] {
					seen[link] = true
					unseenlinks <- depthlink{list.depth, []string{link}}
				}
			}
		} else {
			break
		}
	}
	fmt.Println("The crawler finished!")
}
