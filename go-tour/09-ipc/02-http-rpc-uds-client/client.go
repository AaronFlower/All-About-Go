package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("unix", "/tmp/rpc.sock")
	if err != nil {
		log.Fatal(err)
	}

	// Synchronous call
	name := "Joe"
	var reply string
	err = client.Call("Greeter.Greet", &name, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Go '%s' \n ", reply)
}
