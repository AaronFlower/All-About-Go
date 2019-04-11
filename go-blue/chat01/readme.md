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

需要监听 `join, leave, clients` 这些 channel, 那我们只能用 `select poll` 来查询了。

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



## package websocket

[websocket](https://godoc.org/github.com/gorilla/websocket) 文档。

```
import "github.com/gorilla/websocket"
```

Package websocket implements the WebSocket protocol defined in [RFC 6455](http://tools.ietf.org/html/rfc6455)

### Overview

The Conn type represents a WebSocket connection, A server application calls the `Upgrader.Upgrade` method from an HTTP requests handler to get a `*Conn`

```
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }
    ... Use conn to send and receive messages.
}
```

## Tracing Code Design

利用 Tracing 在程序的关键步骤输出信息是一个很好的实践，可以显示的观察我们的程序是以何种方式运行的。在写 Tracing 代码的时候，我们可以使用 TDD (Test-Driven Development) 实践来进行开发。TDD 可以的保证我们的 package 的可用性。

### 利用 TDD 开发一个 Package

在 Go 中，一个 Package 被组织在一个目录中。Go 没用 subpackage 的概念，所以嵌套目录中的 package 和 父目录中的 package 是没有关系的。

在我们的 Chat 程序中，我们所有的文件都在 `main` package 中，因为我们想构建一个可执行工具。而我们的 Tracing Package 不会被直接执行，构建成一个可执行工具的，所以它可以使用其它的 package name.

我们需要考虑 package 的 API (Application Programming Interface) , 我们应该考虑怎样构建一个库，用户易扩展，易用，已及应该暴露那些信息和隐藏那些信息。

目标：
    - Tracing Package 应该易用
    - Unit test 单元测试应该覆盖功能测试
    - 用户可以灵活替换 tracer 的实现

### 接口 Interfaces

避免具体的实现细节，我们通过接口就可以定义 API。

新建 `tracer` 目录，并且新建一个 `tracer.go` 文件，并在文件中加入下面的内容：

```
package tracer

type Tracer interface {
    Trace (...interface{})
}
```

- package 名称和目录名一致也是一个很好的实践。

### Unit tests, 单元测试

上面我们只定义的接口，没有实现是不能直接进行测试的。但是在具体实现之前我们还是先写测试。新建一个 `tracer_test.go` 文件，然后添加下面的代码：

```
package tracer

import (
    "testing"
)

func TestNew(t *testing.T) {
    t.Error("We haven't written our test yet")
}
```

执行下 `go test` 看下效果。下面我们我编译对 `New` 方法进行测试的代码。

```
func TestNew(t *testing.T) {
    var buf bytes.Buffer
    tracer := New(&buf)
    if tracer == nil {
        t.Error("Return from New should not be nil")
    } else {
        tracer.Trace("Hello trace package.")
        if buf.String() != "Hello trace package." {
            t.Errorf("Trace should not write '%s'.", buf.String())
        }
    }
}
```


我们希望可以用 `bytes.Buffer` 来捕获 Tracer 的输出，然后保证捕获的输出是一直的。

### Red-Green testing

在执行 `go test` 时现在是会报错的，因为我们还没有实现 `New` 函数呀。在没有实现之前就进行测试，这是 `red-green` 测试的实践。Red-green 测试建议我们先写一个单元测试，测试失败后(产生一个错误), 然后我们根据错误进行编译代码，编写最少的代码来通过测试，然后新的函数也是重复这个流程。这种方式的关键是确保我们所写的代码是确实有意义的，是解决具体功能的代码。

我们可以把 `go test` 看成一个 `Todo` 列表，一次解决一项。

接下来们实现 New 方法。

```
func New () {}
```

执行 `go test` 会报参数错误，和返回值错误。这只是一个例子，在你熟悉了测试驱动开发后就可以避免这些细节了。


修改第一个错误：

```
func New(w io.Writer) {}
```
这说明我的 `New` 函数接收的是一个满足 `io.Writer` 接口的参数, 即对应的对象应实现 `Write` 方法。通过 io.Writer 来接收的话，这就可以让用户来决定将内容记录到那里了。输出的目的可以是标准输出，文件，网络 socket, `bytes.Buffer` 以及自定义的对象等。

在加上对应的返回值，我们的函数定义如下：

```
func New(w io.Writer) Tracer{
   return nil
}
```

可以执行 `go test -cover` 来进行下覆盖测试。

### 接口实现 Implementing the interface

为了满足测试，我们需要从 New 方法中返回具体的对象了，毕竟 Tracer 只是一个接口而已。让我加上一个具体的实现吧。在 `tracer.go` 中加入下面的代码：

```
type tracer struct {
    out io.writer
}

func (t *tracer) Trace(a ...interface{}) {}
```

然后 New 函数返回一个 Tracer 对象。

```
func New(w io.Writer) Tracer {
    return &tracer{out: w}
}
```

执行 `go test` 还是不能通过，我们还需要实现具体的方法：

```
func (t *tracer) Trace(a ...interface{}) {
    fmt.Fprint(t.out, a...)
    fmt.Fprintln(t.out)
}
```

现在执行 `go test -cover` 就可以了。

### 向用户隐藏实现

注意我们的`tracer` 并没有导出，用户只需要了解满足 Tracer 接口的对象是什么，并不需要知道我们 tracer 的具体的实现。这也使得我们包的 API 尽可能的简单干净。

### 在 Chat 中使用 tarce 包

在 room struct 中加下人 tracer 字段来进行记录。

### 可选的 trace

当我们的应用程序发布后，我们并不想 tracing 信息还会输出，我们需要抑制输出信息。为了解决这个问题，我们的 `trace` 包可以提供一个 `Off()` 方法，返回一个满足 `Tracer` 接口的对象，但是当 `Trace` 方法调用时不会输出任何信息。

下面先写 Off 的测试用例。

```
fun TestOff(t *testing.T) {
    var silentTracer Tracer = Off()
    silentTracer.Trace("something")
}
```

为了让测试通过，我们需要实现一个当调用 `Trace` 是什么也不做的 tracer. 代码如下：

```
type nilTracer struct {}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
    return &nilTracer{}
}
```

然后在我们 `newRoom` 方法默认创建的 roomm 使用 `Off` 返回的 Tracer 就可以了。


### 干净的 API (Clean Package API)

最终我们的 `trace` 包的提供的一个干净的 API 是：

- New() : 创建一个 Tracer 实例。
- Off() : 创建一个抑制消息的 Tracer.
- Tracer 接口：描述一个 Tracer 应该实现的方法.
