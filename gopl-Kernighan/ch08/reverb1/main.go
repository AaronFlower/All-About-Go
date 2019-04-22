package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	// io.Copy copies from src to dst until either EOF is reached  on src or an error occurs.
	io.Copy(c, c)
	// when EOF is reached, or error occurs we close the connection.
	c.Close()
}
