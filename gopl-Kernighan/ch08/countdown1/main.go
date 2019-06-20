package main

import (
	"fmt"
	"time"
)

// Countdown implements the countdown for a rocket launch
func main() {
	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Life off!")
}
