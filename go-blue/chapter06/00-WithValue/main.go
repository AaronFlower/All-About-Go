package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	type myContextKey string

	f := func(ctx context.Context, k myContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := myContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	// ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, myContextKey("color"))
}
