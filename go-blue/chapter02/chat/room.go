package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	join    chan *client
	leave   chan *client
	forward chan []byte
	clients map[*client]bool
}

const (
	bufferSize = 1024
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  bufferSize,
		WriteBufferSize: bufferSize,
	}
)

func newRoom() *room {
	return &room{
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan []byte),
		clients: make(map[*client]bool),
	}
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
	}

	c := &client{
		conn: conn,
		msg:  make(chan []byte),
		room: r,
	}
	defer func() { r.leave <- c }()
	r.join <- c
	go c.write()
	c.read()
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			client.Close()
			delete(r.clients, client)
		case msg := <-r.forward:
			for client := range r.clients {
				client.msg <- msg
			}
		}
	}
}
