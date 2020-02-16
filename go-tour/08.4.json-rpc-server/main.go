package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"
)

// Args for http rpc request parameters definitions.
type Args struct {
	A, B int
}

// Quotient provides the Divide func definition.
type Quotient struct {
	Quo, Rem int
}

// Arith redefines int.
type Arith int

// Multiply returns the product of two numbers.
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide returns two numbers quotient.
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	// Arith has implements Multiply, Divide methods, and we can use arith to register a rpc
	arith := new(Arith)
	rpc.Register(arith)

	// make rpc to use TCP
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("TCP RPC has been listening at :1234")

	for {
		conn, err := listener.Accept()
		fmt.Println("Client Established: ", time.Now().String())
		if err != nil {
			continue
		}
		// 如果只用  rpc.ServeConn(conn) 的话，它是一个阻塞型的单用户程序，如查想实现并发前面加上 go 就行了。
		// the only difference between jsonrpc and tcp/http rpc.
		// json-rpc based on tcp don't support http.
		go jsonrpc.ServeConn(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
