package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	tick := time.Tick(time.Second * 1)

	count := 10

	fmt.Println("Commencing countdown, Press return to abort.")
	for i := 10; i > 0; i-- {
		select {
		case <-abort:
			fmt.Println("Abort...")
			return
		case <-tick:
			count--
			fmt.Println(count)
		}
	}
	fmt.Println("Life off!")

}
