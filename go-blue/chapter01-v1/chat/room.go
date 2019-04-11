package main

import (
	"log"
	"net/http"

	"github.com/AaronFlower/All-About-Go/go-blue/chapter01-v1/trace"
	"github.com/gorilla/websocket"
)

// Room defines a chat room
type Room struct {
	clients map[*Client]bool
	join    chan *Client
	leave   chan *Client
	forward chan string
	tracer  trace.Tracer
}

// Run starts the chat room
func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.tracer.Trace("[+] A Client is joined.")
		case client := <-r.leave:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.forward) // 通知 client 的 Write 不要再等了.
				r.tracer.Trace("[+] A Client has left.")
			} else {
				r.tracer.Trace("[-] The leaving client is unkonwn")
			}
		case msg := <-r.forward:
			for client := range r.clients {
				client.forward <- msg
				r.tracer.Trace(" -- send to client ")
			}
		}
	}
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	bufferedSize := 1024
	upgrader := websocket.Upgrader{
		ReadBufferSize:  bufferedSize,
		WriteBufferSize: bufferedSize,
	}

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	client := &Client{
		room:    r,
		conn:    conn,
		forward: make(chan string),
	}

	r.join <- client
	defer func() { r.leave <- client }()

	// 等待向浏览器写入信息
	go client.Write()

	// 建立一个长链接，等待读取浏览器信息。
	client.Read()
}

// NewRoom creates a new room to serve.
func NewRoom() *Room {
	return &Room{
		clients: make(map[*Client]bool),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		forward: make(chan string),
		tracer:  trace.Off(),
	}
}
