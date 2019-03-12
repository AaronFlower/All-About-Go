package mipc

import "testing"

type echoServer struct {
}

func (server *echoServer) Handle(method, params string) *Response {
	var body string
	if method == "get" {
		body = "ECHO:" + params
	}
	return &Response{"200", body}
}

func (server *echoServer) Name() string {
	return "EchoServer"
}

func TestIPC(t *testing.T) {
	server := NewIPCServer(&echoServer{})

	client1 := NewIPCClient(server)
	client2 := NewIPCClient(server)

	resp1, _ := client1.Call("get", "From Client1")
	resp2, _ := client2.Call("get", "From Client2")

	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client1" {
		t.Error("IPCClient.Call failed. resp1:", resp1, " resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
