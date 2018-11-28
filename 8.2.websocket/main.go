package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Cant receive")
			break
		}

		fmt.Println("Recieved back from client: " + reply)
		msg := "Received: " + reply
		fmt.Println("Sending to client:", msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cant send")
			break
		}
	}
}
func main() {
	http.Handle("/", websocket.Handler(echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
