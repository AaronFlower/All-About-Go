package main

import "fmt"

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for i := 0; ; i++ {
			numbers <- i
		}
	}()

	// squarer
	go func() {
		for {
			x := <-numbers
			squares <- x * x
		}
	}()

	// printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
