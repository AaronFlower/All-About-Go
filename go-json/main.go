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

	fmt.Println(string(bytes), len(bytes), cap(bytes))

	in := `{"first_name": "John", "last_name": "Dow"}`

	bytes, err = json.Marshal(in)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes), len(bytes), cap(bytes))

	testRawMessge()

}

func testRawMessge() {
	fmt.Println("Test RawMessage ")
	in := `{"first_name": "John", "last_name": "Dow"}`
	rawIn := json.RawMessage(in)

	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	var p Person

	err = json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("p = %+v\n", p)

}
