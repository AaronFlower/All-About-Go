## HTTP and RPC Protocols over UDS

高层的协议，如 HTTP 和其它形式的 RPC 并不关心底层的栈和具体实现。

Network protocols compose by design. High-level protocols, such as HTTP and various forms of RPC, don't particularly care about how the lower levels of the stack are implemented as long as certain guarantees are maintained.

Go 的标准库提供了一个 `rpc` 库可以让我们很容易的实现一个 RPC 服务器和客户端。下面是一个简单的 server 只有一个 procedure.

```
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

const sockAddr = "/tmp/rpc.sock"

type greeter struct {
}

func (g greeter) greet(name *string, reply *string) error {
	*reply = "Hello " + *name
	return nil
}

func main() {
	fmt.Println("vim-go")
	if err := os.RemoveAll(sockAddr); err != nil {
		log.Fatal(err)
	}

	g := new(greeter)
	// 将 g receiver 实现的方法注册到默认的 RPC Server 上。
	rpc.Register(g)
	// 将 HTTP handler 注册到默认的 RPC Server 上.
	rpc.HandleHTTP()

	l, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serving...")
	http.Serve(l, nil)
}
```

该 RPC 服务器是构建在 HTTP 之上的，它注册了一个 HTTP handler ，所以实际处理服务请求的是 `http.Serve`。该服务器的网络栈如下：

     --------------------------
    |       RPC                 |
     --------------------------
    |       HTTP                |
     --------------------------
    |   Unix Domain Socket      |
     --------------------------
