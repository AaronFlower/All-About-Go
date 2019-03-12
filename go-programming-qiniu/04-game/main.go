package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("vim-go")

	// 为什么是 12 个那?
	fmt.Println(runtime.NumCPU())
}
