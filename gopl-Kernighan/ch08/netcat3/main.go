package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("The server lost!")
		}
		done <- struct{}{}
	}()
	go mustCopy(conn, os.Stdin)
	<-done
	fmt.Println("Connection closed!")
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("The sever has been closed.")
		log.Fatal(err)
	}
}
