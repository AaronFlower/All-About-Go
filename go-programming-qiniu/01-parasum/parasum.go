package main

import "fmt"

func sum(data []int, resultChan chan<- int) {
	sum := 0
	for _, v := range data {
		sum += v
	}
	fmt.Printf("%v = %d \n", data, sum)
	resultChan <- sum
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)

	// 跟实际计算的顺序可能不太一们哟。
	go sum(data[:len(data)/2], resultChan)
	go sum(data[len(data)/2:], resultChan)

	sum1, sum2 := <-resultChan, <-resultChan
	fmt.Println(sum1, sum2, sum1+sum2)
}
