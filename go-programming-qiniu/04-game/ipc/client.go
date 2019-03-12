package mipc

import "encoding/json"

// IPCClient defines the client struct to connect the server.
type IPCClient struct {
	conn chan string
}

// NewIPCClient creates a connection to the server.
func NewIPCClient(server *IPCServer) *IPCClient {
	c := server.Connect()
	return &IPCClient{c}

}

// Call send request messages to the server.
func (client *IPCClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}

	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn // wait for the response

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1

	return
}

// Close closes the client connection to the server.
func (client *IPCClient) Close() {
	client.conn <- "CLOSE"
}
