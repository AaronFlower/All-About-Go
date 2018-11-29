package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("The server is listening at ", service)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	fmt.Println("Client established:", daytime)
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}
}
