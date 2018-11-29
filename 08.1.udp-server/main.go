package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkErr(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)

	fmt.Println("The UDP server is listening: ", service)

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	fmt.Println("The client established: ", daytime)
	conn.WriteToUDP([]byte(daytime), addr)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		os.Exit(1)
	}
}
