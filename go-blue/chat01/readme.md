## ChatRoom 建模

一个 ChatRoom 应用会将所有用户都加入到一个聊天室内，这个聊天室负责 client 的链接，以及负责消息的路由接收和广播。

那么们首先需要创建一个 room, 而用 client 代表一个连接。


### client 建模

一个 client 至少要有的三个属性：
1. 当前的 websocket 链接, `conn`
2. 所属于的聊天室, `*room`
3. 用于接收聊天室的广播信息的 buffered channel, `send`

client 所要做的事：
1. 监听读取浏览器用户发送过来的数据，然后存放到 room 的消息队列上。 `read()/readFromSocket`
2. 监听 room 广播的数据，然后发送给浏览器用户。`write()/writeToSocket`

代码如下:

```go
type client struct {
    socket *websocket.Conn
    room   *room
    send   chan []byte
}

func (c *client) read() {
    // read from socket
    // push to room message queue.
}

func (c *client) write() {
    // wait and read from room
    // send to websocket
}
```

### room 建模

一个 room 需要管理 client 的 join 和 leave, 也需要管理消息队列的转发，所以一个 room 包含的属性：
1. 接收 client 消息的消息队列 buffered channel, `forward / messages`
2. 等待加入的 client channel, `join`
3. 等待离开的 client channel , `leave`
4. 管理所有 clients 的 map 表, `clients`

room 所要做的事：
1. run() 开启一个聊天室，负责 client 的加入删除和消息转发。
2. 实现对 http 请求的处理，基为我们的浏览器是基于 http 的。
```go
type room struct {
	forward chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func (r *room) run() {

}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

```

### Turning a room into an HTTP handler


channel 可以看成是一个内存级线程安全的消息队列，发送者存数据而接收取数据。在聊天室 room 的 channel 是提供给 client 放数据的地方，放了数据之后 room 会把这数据取出广播到 client 的 channel 上，client得知数据已经被存上了之后，会把数据取走发给客户端浏览器。


## Tracing Code Design
