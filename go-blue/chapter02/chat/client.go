package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a single chatting user.
type Client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn

	// send is a channel on which message are sent.
	send chan []byte

	// room is the room this client is chatting in.
	room *Room
}

// Read allows client to read message from the socket
// via the ReadMessage method.
// And send any received messages to the forward channel
// on the room type.
func (c *Client) Read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		c.room.forward <- msg
	}
}

// Write accepts message from the send channel.
func (c *Client) Write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
