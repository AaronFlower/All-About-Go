package main

import "encoding/json"
import "fmt"

// Person The person
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Message the message
type Message struct {
	Name string
	Body string
	Time int64
}

// FamilyMember for references test.
type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func testMarshal() {
	fmt.Println("Test Marshal")
	m := Message{"Eason", "Hello", 231231231}

	bytes, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes), len(bytes), cap(bytes))
}

func testUnmarshalReferences() {
	fmt.Println("\n ***** Test Unmarshal Refercences \n")
	b := []byte(`{"Name": "Eason", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var m FamilyMember
	err := json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("m = %+v\n", m)
}

func testUnmarshal() {
	fmt.Println("Test Unmarshal")
	in := Message{Name: "Aaron", Body: "This is a  body"}
	b, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	var m Message
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("m = %+v\n", m)
}

func testDecodingArbitraryData() {
	fmt.Println("")
	fmt.Println("Test Decoding arbitrary data")
	fmt.Println("")
	b := []byte(`{"Name": "Wendnesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("f = %+v\n", f)

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
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

	testMarshal()

	testUnmarshal()

	testRawMessge()

	testDecodingArbitraryData()

	testUnmarshalReferences()

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
