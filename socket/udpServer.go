package socket

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient(conn *net.UDPConn) {
	for {
		reader := bufio.NewReader(conn)
		buf := make([]byte, 1024)
		_, err := reader.Read(buf)
		if err != nil {
			continue
		}
		fmt.Println(string(buf))
	}
	//dayTime := time.Now().String()
	//conn.WriteToUDP([]byte(dayTime), addr)
}

func UDPServer() {
	//service := ":3000"
	//udpAddr, err := net.ResolveUDPAddr("udp", service)
	//checkError(err)
	var err error
	conn, err = net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 3000,
	})
	checkError(err)
	fmt.Println("udp server")
	for {
		handleClient(conn)
	}
}
