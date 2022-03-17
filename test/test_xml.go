package test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type Servers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
}

type server struct {
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func TestDecXml(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err !=nil{
		fmt.Printf("error: %v", err)
		return
	}

	v :=Recurlservers{}
	if err:=xml.Unmarshal(data,&v);err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}

func TestEncXml(){
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
