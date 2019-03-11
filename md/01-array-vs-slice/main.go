package main

import "fmt"

func walk(values []int) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("vim-go")
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}

	walk(b)
	fmt.Printf("b = %+v, len = %d, cap = %d \n", b, len(b), cap(b))
	fmt.Printf("a = %+v, len = %d, cap = %d \n", a, len(a), cap(a))
}
