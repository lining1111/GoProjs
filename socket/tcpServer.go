package socket

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func handlerClient(conn net.Conn) {
	defer conn.Close()
	dayTime := time.Now().String()
	conn.Write([]byte(dayTime))

}

//func init() {
//	service := ":7777"
//	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
//	checkError(err)
//	listener, err := net.ListenTCP("tcp", tcpAddr)
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			continue
//		}
//		go handlerClient(conn)
//	}
//}
