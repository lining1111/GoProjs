package socket

import (
	"fmt"
	"net"
	"os"
)

func init() {
	service := ":7777"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	_,err = conn.Write([]byte("anything"))
	checkError(err)
	var buf []byte
	n,err :=conn.Read(buf)
	checkError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
