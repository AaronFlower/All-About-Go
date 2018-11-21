## How to correctly Serialize JSON String in Golang

### Demo 01
```
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
1. [How To Correctly Serialize JSON String In Golang](http://goinbigdata.com/how-to-correctly-serialize-json-string-in-golang)
