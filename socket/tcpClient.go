package socket

import (
	"io/ioutil"
	"net"
)

var client *net.TCPConn

//func init() {
//	service := "127.0.0.1:3000"
//	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
//	checkError(err)
//	conn, err := net.DialTCP("tcp", nil, tcpAddr)
//	checkError(err)
//	//_, err = conn.Write([]byte("HEAD / HTTP/1.0/\r\n\r\n"))
//	//checkError(err)
//	client = conn
//	result, err := ioutil.ReadAll(conn)
//	checkError(err)
//	fmt.Println(string(result))
//	os.Exit(0)
//
//}

func GetTCPInfo() []byte {
	result, err := ioutil.ReadAll(client)
	if err != nil {
		return []byte(err.Error())
	}
	return result
}
