package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Server defines a server struct.
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

// Serverslice stores a bucket of servers
type Serverslice struct {
	Servers []Server `json:"servers"`
}

func parseUnknownJSON() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}

	if m, ok := f.(map[string]interface{}); ok {
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is array")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a unknown type.")
			}
		}
	}
}

func main() {
	file, err := os.Open("servers.json")
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}

	var s Serverslice
	json.Unmarshal(data, &s)
	fmt.Println(s)

	fmt.Println("Parse unknown JSON")
	parseUnknownJSON()
}
