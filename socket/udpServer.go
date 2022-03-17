package socket

import (
	"net"
	"time"
)

func handleClient(conn *net.UDPConn) {
	var buf []byte
	_, addr, err := conn.ReadFromUDP(buf)
	checkError(err)
	dayTime := time.Now().String()
	conn.WriteToUDP([]byte(dayTime), addr)
}

func init() {
	service := "1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}
