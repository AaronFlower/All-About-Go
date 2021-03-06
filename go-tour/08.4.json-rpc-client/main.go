package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
	"strconv"
)

// Args provides rpc function communicate parameters definition.
type Args struct {
	A, B int
}

// Quotient receives Divide method result.
type Quotient struct {
	Quo, Rem int
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: ", os.Args[0], "server:port a b")
		os.Exit(1)
	}
	service := os.Args[1]

	// client, err := rpc.Dial("tcp", service)
	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing", err)
	}

	// Synchronous call
	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Usage: ", os.Args[0], "server:port a b")
		os.Exit(1)
	}
	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Usage: ", os.Args[0], "server:port a b")
		os.Exit(1)
	}

	args := Args{a, b}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d / %d = %d remainder %d \n", args.A, args.B, quot.Quo, quot.Rem)

}
