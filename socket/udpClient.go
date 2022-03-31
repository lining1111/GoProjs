package socket

import (
	"fmt"
	"net"
)

var conn *net.UDPConn

func UDPConnect() {
	service := ":3000"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)
	conn, err = net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	fmt.Println("udp connect")
	//_,err = conn.Write([]byte("anything"))
	//checkError(err)
	//var buf []byte
	//n,err :=conn.Read(buf)
	//checkError(err)
	//fmt.Println(string(buf[0:n]))
	//os.Exit(0)
}

func GetUDPInfo() []byte {
	var buf []byte
	n, err := conn.Read(buf)
	if err != nil {
		return []byte(err.Error())
	}
	if n > 0 {
		return buf[0:n]
	}

	return nil
}
