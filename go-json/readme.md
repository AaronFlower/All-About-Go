## How to correctly Serialize JSON String in Golang

### JSON And GO

使用 `json` 包可以很方便的与 JSON 打交道。

### 1. Encoding 编码

使用 `Marshal` 函数来进行编码。函数原型如下：

```go
func Marshal(v interface{}) ([]byte, error)
```

函数接受任何类型。

定义下面的类型:

```go
type Message struct {
    Name string
    Body string
    Time int64
}

m := Message{"Eason", "Hello", 231231231}
b, err := json.Marshal(m)
fmt.Println(string(b)) // {"Name":"Eason","Body":"Hello","Time":231231231}
```

### 2. Decoding 解码

使用 `Unmarshal` 来进行解码。可以将 byte 数据流转换成任何数据，函数原型为:

```go
func Unmarshal(data []byte, v interface{}) error
```

```go
in := `{"Name": "Aaron", "Body": "bye"}`
b := json.Marshal(in)
var m Message
err := json.Unmarshal(b, &m)
```

### 3. 解码任意数据

对于事先不知道的数据结构类型，我们可以将其解决到 `interface{}`. 

```go
func testDecodingArbitraryData() {
    b := []byte(`{"Name": "Wendnesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
    var f interface{}
    err := json.Unmarshal(b, &f)
    if err != nil {
        panic(err)
    }
    fmt.Printf("f = %+v\n", f)
}
```

f 会被解码成一个 `f map[string]interface{}`.

```go
f = map[Name:Wendnesday Age:6 Parents:[Gomez Morticia]]
```

我们可以通过 `range` 来遍历这个 map, 并且使用 `type switch`来获取正确的值和类型。

### 4. 解决引用数据



### Demo 01

```go
package main

import "encoding/json"
import "fmt"

// Person
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	bytes, err := json.Marshal(Person{
		FirstName: "John",
		LastName:  "Dow",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes), len(bytes))

	in := `{"first_name": "John", "last_name": "Dow"}`

	bytes, err = json.Marshal(in)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
```
#### Output
```
{"first_name":"John","last_name":"Dow"} 39 48
"{\"first_name\": \"John\", \"last_name\": \"Dow\"}" 52 64
```
Can you see escaped double quotes? This is the crux of the problem. `json.Marshal` escapes string while serializing it.

#### Skipping 转义(skipping escaping)

`json` 包提供了一个解决文案。使用 `RawMessage` 类型，使用该类型进行 `Marshal, Unmarshal` 时不会进行转义。

### References

1. [JSON and go](https://blog.golang.org/json-and-go)
2. [How To Correctly Serialize JSON String In Golang](http://goinbigdata.com/how-to-correctly-serialize-json-string-in-golang)