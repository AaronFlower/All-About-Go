package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client defines a client.
type Client struct {
	conn    *websocket.Conn
	room    *Room
	forward chan string
}

// Read reads data from the browser and send it the room
func (c *Client) Read() {
	defer c.conn.Close()
	for {
		_, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			c.room.leave <- c
			break
		}
		c.room.forward <- string(p)
	}
}

// Write writes the rooms forcasting message to the browser.
func (c *Client) Write() {
	defer c.conn.Close()
	for msg := range c.forward {
		c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}
