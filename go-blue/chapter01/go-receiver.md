## Go Receiver 

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

在 `var i I = T{"hello"}` 中，我们是直接用 T 给 i 赋值。但是如果在实现接口函数时用的是指针接收的话，那么就需要用下面的方式来初始化 i 了, 即 `var i I = &T{"hello"}`。

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = &T{"hello"}
	i.M()
}
```

这一点还是要注意的。



