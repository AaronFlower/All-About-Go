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

// 这个爬虫程序有一个问题， too parallel. 无限制的并行通常是不太可取的，因为我们的资源是有限的。
// 所以在下一个例子中，我们要限制一下其并行度。The program is too parallel.
func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	go func() {
		worklist <- os.Args[1:]
		fmt.Println("os.Args[1:]", os.Args[1:])
	}()

	// crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		fmt.Println("list:", list)
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
