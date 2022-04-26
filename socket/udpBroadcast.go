package socket

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

var connBroadcast *net.UDPConn
var localPort int
var Num int

func Open(port int) error {
	localPort = port
	var err error
	connBroadcast, err = net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: port,
	})
	return err
}

func broacastContent(content []byte) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()
	connBroadcast.Write(content)
}

func BroadCast(content []byte) {
	if connBroadcast != nil {
		go func() {
			for true {
				time.Sleep(time.Duration(10) * time.Second) //10s sleep
				_, err := connBroadcast.Write(content)
				if err != nil || Num == 2 {
					//重新打开
					fmt.Println("重新打开")
					//Close()
					Open(localPort)
				}
			}
			fmt.Println("广播退出")
		}()

	}
}

func Close() {
	conn.Close()
}
