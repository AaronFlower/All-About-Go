package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	msg  chan []byte
	conn *websocket.Conn
	room *room
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) read() {
	defer c.conn.Close()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		c.room.forward <- msg
	}
}

func (c *client) write() {
	defer c.conn.Close()
	for msg := range c.msg {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
