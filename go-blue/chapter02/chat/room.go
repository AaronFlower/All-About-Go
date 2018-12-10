package main

import (
	"log"
	"net/http"

	"github.com/aaronflower/ago/go-blue/chapter01/trace"
	"github.com/gorilla/websocket"
)

// Room provides a chatting room.
type Room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte

	// join is a channel for clients wishing to join the room.
	join chan *Client

	// leave is a channel for clients wishing to leave the room.
	leave chan *Client

	// clients holds all current clients in this room.
	clients map[*Client]bool

	// tracer will receive trace information of activity in the room.
	tracer trace.Tracer
}

// NewRoom makes a new room.
func NewRoom() *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		tracer:  trace.Off(),
	}
}

// Run starts the chat Room
func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
			r.tracer.Trace("New client joined")
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
			r.tracer.Trace("Client left")
		case msg := <-r.forward:
			// forward message to all clients
			r.tracer.Trace("Message received: ", string(msg))
			for client := range r.clients {
				client.send <- msg
				r.tracer.Trace(" -- send to client")
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBUfferSize = 1024
)

var (
	upgrader = &websocket.Upgrader{
		ReadBufferSize:  socketBufferSize,
		WriteBufferSize: socketBufferSize,
	}
)

// ServeHTTP defines the room to act as a hanlder
func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Upgrade HTTP connection to websocket.
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &Client{
		socket: socket,
		send:   make(chan []byte, messageBUfferSize),
		room:   r,
	}
	r.join <- client
	defer func() { r.leave <- client }()
	go client.Write()
	client.Read()
}
