package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	fmt.Println("vim-go")
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("[-] Error ", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn)

	for input.Scan() {
		go echo(conn, input.Text(), 1*time.Second)
	}
	conn.Close()
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}
