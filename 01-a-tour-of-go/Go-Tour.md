## Go-Tour

### 1. defer 语句

- `defer` 在函数返回时会执行，即使发生 `panic` 也会执行。
- 在遇到`defer` 语句时，defer 中表达式会立即进行求值，但只有在函数返回时才会执行。

```go
package main

import "fmt"
import "errors"

func main() {
	defer fmt.Println("world")
	panic(errors.New("Test panic"))
	fmt.Println("hello")
}

```

### 2. pointer

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var ptr *int
	fmt.Printf("ptr = %v\n", ptr) // ptr = <nil>
	fmt.Printf("ptr = %v\n", &ptr) // ptr = 0xc00000e010
	fmt.Printf("ptr = %v\n", *ptr) // Invalid memory call.
	var pi int
	fmt.Printf("pi = %v\n", &pi)
	fmt.Printf("pi = %v\n", pi)
  
	var v Vertex // 声明时如果未初始化，其 fields 会按默认值来初始化。也会分配地址的。
  
	fmt.Printf("v = %v\n", v)
	fmt.Printf("v = %v\n", &v)
  p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

### 3. Arrays

- 数组的长度是数组类型的一部分，所以是不能 resize 的。
- `var  a [10]int`
- `var b [...]int{1, 2, 6}`, `[...] ` 可用来代替长度。

```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	odds := [...]int{1, 3, 5, 7}
	fmt.Printf("v = %v, %T", odds, odds)
}
```

输出结果:

```go
Hello World
[Hello World]
v = [1 3 5 7], [4]int
Program exited.
```

### 4. Slices

- `[]T` 是一个类型为 `T` 的 slice;
- slice 的长度是动态的;
- slice 可以看作是数组的引用;
- 创建 slice 有三种方式：
  - 通过对数组的 `arr[begin:end]` 创建；
  - 通过 slice literal 创建，如 `p := []int{1, 2}`
  - 通过 `make` 来创建，如: `a :=make([]int, 5) // len(a) = 5`

- slice 具有 `length` 和 `capacity`
  - `len(s)` 获取 slice 和长度，即其实际拥有的数组个数；
  - `cap(s)` 获取其容量，即其基于数组的长度。
  - slice 可以 `re-slicing`, 如：`s = s[2:] // Drop its first two values`

- 无元素的 slice 是 `nil`.

- 我们可以创建 slice 的 slice 即多维 slice

  ```go
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }
  ```

- `func append(s []T, vs ...T) []T` 通过内置函数 `append` 来向一个 slice 添加一个或多个元素。

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	
	fmt.Println(s)
	fmt.Printf("s's type is :%T\n", s) // slice []int 无长度
	fmt.Printf("primes's type is :%T\n", primes) // 数组，[6]int 有长度
  
  q := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("q = %T\n", q)
}
```

输出结果：

```
[3 5 7]
s's type is :[]int
primes's type is :[6]int
```

- Exercise: Slices

  ```go
  package main
  
  import "golang.org/x/tour/pic"
  
  func Pic(dx, dy int) [][]uint8 {
  	pic := make([][]uint8, 0)
  	for x := 0; x < dx; x++ {
  		pic = append(pic, make([]uint8, 0))
  		for y := 0; y < dy; y++ {
  			pic[x] = append(pic[x], uint8((x + y) / 2))
  		}
  	}
  	return pic
  }
  
  func main() {
  	pic.Show(Pic)
  }
  ```

### 5. Maps

- 只声明，不使用 make 初始化的 map 其值为 `nil`.

- 使用 make 来初始化一个 map，然后才能被使用.

- 也可以使用 map literal 来初始化一个 map

  ```go
  var m = map[string]int {
    "Apple": 1,
    "Android": 2,
  }
  ```

### 6. Methods

- 只能为同一个包（pacakge) 中的类型添加方法，而不能为其它包中的类型添加方法。如 built-in 的 int. 

- 如果将 built-in 的 int 添加方法的话，我们需要用 `type` 来声明另一个类型。

  ```go
  type MyFloat float64
  
  func (f MyFloat) Abs() float64 {
  	if f < 0 {
  		return float64(-f)
  	}
  	return float64(f)
  }
  ```

- 指针接收器（Pointer receivers)

  在为类型添加方法时，也可以用 pointer 来接收。一般使用指针来接收更加常见，因为通过指针可修改接收者的值。

- 当定义一个指针接收方法时，在调用时即使我们不使用指针也是可以的。Golang 会自动为我们转换成指针。
- 使用指针接收器的两个原因：
  - 指针传递的是一个引用，避免每次调用的数据 copy;
  - 在调用时即可用变量也可用指针变量会自动转换;

### 7. Interfaces

接口是用来定义一组方法签名的集合。任何类型都可以实现这些方法。

- 注意在用指针类型和非指针类型来实现方法时，在赋值时的方式是不一样的。
- interface 在调用方法时，一定要实例化；
- `interface{}` 是一个空接口。

### 8. Type assertions

```go
t := i.(T) // If i does not hold a T, the statement will trigger a panic.
t, ok := i.(T)
```

另外也可以用 `type` 关键字拉倒 `switch` 来判断：

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

### 9. Stringers

一个最常用到接口是 `Stringer` ，它只定义了一个方法。

```go
type Stringer interface {
  String() string
}
```

`fmt` 包默认会使用 `String` 方法进行输出。

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

### 10. Errors

`error` 类型也是一个常用的内置类型，与 `fmt.Stinger` 接口类似，它也定义一个方法：

```go
type error interface {
  Error() string
}
```

我们可以自定义 Error 通过实现 Error 方法。

```go
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
```

### 11. Readers

Golang 在定义接口时，方法最小化原则。可以通过实现不同的接口来组合。如 `io.Reader` 定义了一个 Read 方法。

```go
func (T) Read(b []byte) (n int, err error)
```

