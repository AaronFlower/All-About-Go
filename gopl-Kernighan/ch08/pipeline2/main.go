package main

import "fmt"

func main() {
	numbers := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	go func() {
		for v := range numbers {
			squares <- v * v
		}
		close(squares)
	}()

	for v := range squares {
		fmt.Println(v)
	}
}
