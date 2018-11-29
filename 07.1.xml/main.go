package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type recurlyServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func generateXML() {
	v := &recurlyServers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"BeiJing_VPN", "127.0.0.1"})
	output, err := xml.MarshalIndent(v, " ", "  ")
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}

	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func main() {
	file, err := os.Open("servers.xml")
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

	v := recurlyServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}
	fmt.Println(v)

	fmt.Println("Generate XML Content")
	generateXML()
}
