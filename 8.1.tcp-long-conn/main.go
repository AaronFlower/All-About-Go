package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("vim-go")
	server := ":2048"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	fmt.Println("The server is listening:", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(10 * time.Second)) // set 10 seconds timeout
	defer conn.Close()
	request := make([]byte, 128) // set maxium request length to 128B to prevent flook attack.

	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Printf("err = %+v\n", err)
			break
		}

		fmt.Println("Client Established, clien say ", string(request[:readLen]), string(time.Now().String()))

		if readLen == 0 {
			break // connection already closed by client.
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte("Your request has been submited."))
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte("Your request has been submited."))
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear last read content.
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		os.Exit(1)
	}
}
