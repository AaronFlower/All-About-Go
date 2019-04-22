package main

import "fmt"

func counter(numbers chan<- int) {
	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)
}

func squarer(numbers <-chan int, squares chan<- int) {
	for v := range numbers {
		squares <- v * v
	}
	close(squares)
}

func printer(squares <-chan int) {
	for v := range squares {
		fmt.Println(v)
	}
}

func main() {
	numbers := make(chan int)
	squares := make(chan int)
	go counter(numbers)
	go squarer(numbers, squares)
	printer(squares)
}
