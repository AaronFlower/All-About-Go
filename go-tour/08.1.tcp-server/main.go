package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)

	fmt.Println("The server is listening:", tcpAddr)

	for {
		conn, err := listener.Accept()
		daytime := time.Now().String()
		fmt.Println("conn established.", daytime)
		if err != nil {
			continue
		}
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		os.Exit(1)
	}
}
