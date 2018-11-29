package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
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
	// make rpc to use HTTP
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
