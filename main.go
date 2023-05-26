package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

//// ReadCmdOut 实时从脚本获取运行结果
//func ReadCmdOut() {
//	shell := "./update.sh"
//	//增加可执行属性
//	cmd := exec.Command("chmod", "+x", shell)
//	cmd.Run()
//	cmd = exec.Command("/bin/bash", "-c", shell)
//	cmd.Stdin = os.Stdin
//	//捕获标准输出
//	stdout, err := cmd.StdoutPipe()
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	if err = cmd.Start(); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	for true {
//		tmp := make([]byte, 1024)
//		_, err := stdout.Read(tmp)
//		fmt.Print(string(tmp))
//		if err != nil {
//			break
//		}
//	}
//	if err = cmd.Wait(); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	return
//}

func handleconnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

func main() {

	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Println(err)
		return
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":8000", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("conn connect ", conn.LocalAddr().String())
		go handleconnection(conn)
	}

	//ReadCmdOut()
	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	socket.Open(3000)
	//	socket.BroadCast([]byte("hello"))
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	for true {
	//		defer func() {
	//			if r := recover(); r != nil {
	//				fmt.Println(r)
	//			}
	//		}()
	//		fmt.Println("请输入数字")
	//		var num int
	//		fmt.Scanf("%d",&num)
	//		socket.Num =num
	//
	//		if num==2 {
	//			panic("got 2")
	//		}else {
	//			fmt.Println(num)
	//		}
	//	}
	//
	//}()
	//wg.Wait()
}
