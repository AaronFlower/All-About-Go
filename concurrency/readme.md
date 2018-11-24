## Go Concurrency

Go 是 21 世纪的语言。

### goroutine

```
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行
}
```

To Test
```
go build

# to run 10 times.

bash -c 'i=0; while [ $i -lt 10 ]; do echo "Exec $i:"; concurrency; i=$[$i + 1];done'

Exec 0:
The concurrency test
HELLO
world
HELLO
HELLO
world
Exec 1:
The concurrency test
HELLO
HELLO
HELLO
Exec 2:
The concurrency test
HELLO
HELLO
HELLO
Exec 3:
The concurrency test
HELLO
world
world
world
HELLO
HELLO
Exec 4:
The concurrency test
HELLO
HELLO
HELLO
world
world
world
Exec 5:
The concurrency test
HELLO
HELLO
HELLO
Exec 6:
The concurrency test
HELLO
HELLO
HELLO
Exec 7:
The concurrency test
HELLO
world
HELLO
world
HELLO
world
Exec 8:
The concurrency test
HELLO
world
world
world
HELLO
HELLO
Exec 9:
The concurrency test
HELLO
HELLO
HELLO
concurrency master
```
