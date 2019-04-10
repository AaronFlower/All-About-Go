## 04-A-Tout-Of-Go : Concurrency

### Goroutines 协程

A goroutine is a lightweight thread managed by the go runtime. 协程是 go 运行时维护的一个轻量级线程。

```go
go f(x, y, z)
```

这就启动一个协程，`f, x, y,z` 的估值在当前协程中发生，`f` 在一个新的协程中执行。

启动的协程是运行在同一个地址空间中的，所以在访问共享内存的时候需要有同步机制才行。`sync` 包提供了一些有用的原语可以用来同步。不过在 Go 中你可能不一定非要使用原语来进行同步，可以使用 channel.

### Channels

Channels 是一个具有类型的通道，你可使用 channel 操作符 `<-` 进行发送和接收数据。

```go
ch <- v 		// send v to channel ch
v := <- ch 	// Receive from ch, and assign value to v.
```

和 map, slices 一样，channels 在使用之前也需要创建.

```go
ch := make(chan int)
```

通常情况下，在一方未准备好之前，send, receive 都会被 block 的，这样就可以使得协程不需要明确的锁或条件变更来进行同步了。

By default, sends and receives block until the other side is ready. This allow goroutines to synchronize without explicit locks or condition variables.

### Buffered Channels

Channels 可以设置成缓冲类型的, 在 make 函数中提供第二个参数就行了。

```go
ch := make(chan int, 100)
```

在发送时，如果 channels 已经满了就会被 block, 面在 receive 时如果 channels 是空的也会被 block.

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

### Range and close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```go
v, ok := <-ch
```

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop `for i := range c` receives values from the channel repeatedly until it is closed.

**Note**: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic

**Another Note**: Channels aren't like files: you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

Channels 和文件不一样，通常是不需要关闭的。只有在需要通知接收方没有消息发送时才关闭，如结束接收方的 `for range`.

### Select

The `select` statement lets a goroutine wait on multile communication operations.

A `select` blocks until one of its cases can run, then it executes that case. It choose one at random if multiple are ready.

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // 向 c 发送数据，在 c 满的时候未被取走会被 block. 
			x, y = y, x+y
		case <-quit: // 接收 quit 数据, 在未写入时也会被 block.
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

### Default Selection

在 `select` 中的 `default` case 在其它 case 被 block 时执行，可以用 default 来完成超时操作。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond) // Tick 函数返回的是 `<-chan Time` 类型，只能接收
	boom := time.After(500 * time.Millisecond) // After 函数返回的是 `<-chan Time` 类型，只能接收
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

从上面的例子中，我们可以函数和函数返回值中定义  channel 的方向，定义是否只能接收或发送。

### sync.Mutex

通过 channels 我们可以在协程间协程通信，但是如果我们希望在协程之间互斥怎么操作那？Go 的标准库 `sync.Mutex` 提供了互斥锁来解决这个问题。它提供两个方法：

- Lock
- Unlock

在上锁之后，可以用 `defer` 进程释放。避免出现死锁的情况。











