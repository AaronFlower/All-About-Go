// "\r" overrides the last character
package main

import (
	"fmt"
	"time"
)

func fib(x int) int {
	if x < 2 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r %c ", r) // "\r" override the last character
			time.Sleep(delay)
		}
	}
}

func main() {
	fmt.Println("vim-go")
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\r Fibonacci(%d) = %d \n", n, fibN)
}
