package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Client represents a single chatting user.
type Client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn

	// send is a channel on which message are sent.
	send chan *message

	// room is the room this client is chatting in.
	room *Room

	// userData holds information about the user
	userData map[string]interface{}
}

// Read allows client to read message from the socket
// via the ReadMessage method.
// And send any received messages to the forward channel
// on the room type.
func (c *Client) Read() {
	defer c.socket.Close()
	for {
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		if avatarURL, ok := c.userData["avatar_url"]; ok {
			msg.AvatarURL = avatarURL.(string)
		}
		c.room.forward <- msg
	}
}

// Write accepts message from the send channel.
func (c *Client) Write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
