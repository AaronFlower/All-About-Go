package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
}

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gofunc"
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
