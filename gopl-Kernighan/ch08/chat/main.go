package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	join  = make(chan client)
	leave = make(chan client)
	msg   = make(chan string)
)

func startChatRoom() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-join:
			clients[cli] = true
			fmt.Println("have join")
		case cli := <-leave:
			fmt.Println("have left")
			delete(clients, cli)
			close(cli)
		case m := <-msg:
			fmt.Println("broadcast")
			for cli := range clients {
				cli <- m
			}
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWrite(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	msg <- who + " has arrived"
	join <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg <- who + ":" + input.Text()
		fmt.Println("scan msg...")
	}

	leave <- ch
	msg <- who + " has left"
	conn.Close()
}

func clientWrite(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	go startChatRoom()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
