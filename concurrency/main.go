package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main() {
	go say("world") // to start a Goroutine to execute.
	say("HELLO")
}
