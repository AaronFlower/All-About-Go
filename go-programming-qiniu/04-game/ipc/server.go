package mipc

import (
	"encoding/json"
	"fmt"
)

// IPC 进程间通信, 简单的 IPC. 这个库的目的简单，就是封装通信包的细节，让使用者可以专注于业务。
// 该库： 用 channel 作为模块之间的通信方式。

// Request defines the Request struct
type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

// Response defines the Response struct
type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

// Server defines a server interface
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

// IPCServer defines an ipc server.
type IPCServer struct {
	Server
}

// NewIPCServer creates an IPC server instance.
func NewIPCServer(server Server) *IPCServer {
	return &IPCServer{server}
}

// Connect returns a connection to server.
func (server *IPCServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				// close the connection
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format: ", request)
			}

			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b)
		}
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}
