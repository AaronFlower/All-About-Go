package main

// Crawl2 crawls web links starting with command-line arguments.
//
// 利用 buffered channel 作为信号量
// This version uses a buffered channel as a counting semaphore.
// to limit the number of concurrent calls to links.Extract.
import (
	"fmt"
	"log"
	"os"

	"github.com/AaronFlower/All-About-Go/gopl-Kernighan/ch05/links"
)

// 利用 buffered channel 来作为一个统计信号量
// tokens is a counting semaphore used to enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	// acquire a token
	tokens <- struct{}{}
	list, err := links.Extract(url)
	// relese a token
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return list
}

// 注意 worklist 是一个 channel ，因为在递归的过程中我需要不断接收新的 worklist
// 而不仅仅是初始化时用户的输出。
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// start with the command-line arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
