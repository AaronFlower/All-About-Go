## IPC 框架

一个简单的 IPC 框架。IPC ? Interprocess communication (IPC) , 什么是一个 IPC 服务器那？IPC 服务器是一台机器的概念。

Interprocess communication (IPC)  allows bidirectional communication between clients and servers using distributed applications. IPC is a mechanism used by programs and multi-user processes. IPCs allow concurrently running tasks to communicate between themselves on a local computer or between the local computer and a remote computer.

### 1. 定义数据结构

- Request : 客户端 client 的请求数据。

```go
type Request struct {
    Method string
    Params string
}
```


- Response : 服务器的响应数据.

```go
type Response struct {
    Code    string
    Body    string
}
```

### 2. 定义一个服务器接口

我们定义的服务器只有简单的两个接口，即 `Name()` 提供服务器的信息，而 `Handle()` 用来处理请求。因为我们的服务器可能有 WEB 服务器，长连接服务器，邮件服务器等，而各个服务器有些特写的操作就不需要定义在接口中，我们接口只定义最简单的基础的接口方法，特定的接口由特写的服务器实现时再来自己定义。


```go
type Server interface {
    Name () string
    Handle (method, request string) *Response
}
```

### 3. 定义 IPCServer

我们现在定义一个IPCServer, 该服务器除了负责实现 Server 接口定义的方法外，而且还应该负责 client 的连接。那么 Server 与 Client 怎么通信那？可以通过 Channel 来实现通信。

任何实现了 Server 接口的服务器都要来初始化 IPCServer。这个可以用嵌入组合实现。

另外我们还提供一个工厂函数，创建一个 IPCServer 实例。

```go
type IPCServer struct {
    Server
}

// Connect 由服务器创建一个客户端连接，并且一直监听这个连接，来处理客户端的语法。该函数返回一个 channel。服务器与客户端之间通过个 Channel 来实现通信。
func (IPCServer *server)Connect() chan string


//

```

### 4. 定义客户端

我们的客户端通过一个 channel 来实现服务器通信，那么我们的客户端的结构体的定义可以只有一个属性就了。至于客户端的工厂函数是需要通过服务器来定义的，所在工厂函数中通过服务器的 `Connection` 函数来实现即可。

客户端仅有两个方法，那就是用来请求服务器的 `Call` 方法，以及主动请求 `Close` 关闭链接的方法。这个方法的通信都需要使用服务器返回的 `channel` 来实现。

```go
type IPCClient struct {
    conn chan string
}

func NewIPCClient(server *IPCServer) *IPCClient {

}

func (client *IPCClient)Call(method, params string)(resp *Response, err error) {
}

func (client *IPCClient)Close() {
    client.conn <- "CLOSE"
}
```

## 中央服务器

一个中央服务器作为全局唯一的实例，从原则上需要承担以下责任：

- 在线玩家的状态管理
- 服务器管理
- 服天系统

中内服务器可管理其它服务器。但是目前我们仅实现了一个服务器那就可先空着了。目前仅提供聊天服务器，聊天系统只实广播。私聊的需求可自己扩展。
