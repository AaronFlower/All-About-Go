package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

const sockAddr = "/tmp/rpc.sock"

// Greeter defines an exported struct for rpc test
type Greeter struct {
}

// Greet greets the people
func (g Greeter) Greet(name *string, reply *string) error {
	*reply = "Hello " + *name
	return nil
}

func main() {
	fmt.Println("vim-go")
	if err := os.RemoveAll(sockAddr); err != nil {
		log.Fatal(err)
	}

	g := new(Greeter)
	// 将 g receiver 实现的方法注册到默认的 RPC Server 上。
	rpc.Register(g)
	// 将 HTTP handler 注册到默认的 RPC Server 上.
	rpc.HandleHTTP()

	l, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serving...")
	http.Serve(l, nil)
}
