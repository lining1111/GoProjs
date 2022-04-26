package main

import (
	"fmt"
	"os"
	"os/exec"

	//"text/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func process(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("form parse fail")
	}
	first_name := r.PostForm["first_name"]
	last_name := r.PostForm["last_name"]
	fmt.Printf("first_name:%s\n", first_name)
	fmt.Printf("last_name:%s\n", last_name)
}

// ReadCmdOut 实时从脚本获取运行结果
func ReadCmdOut() {
	shell := "./update.sh"
	//增加可执行属性
	cmd := exec.Command("chmod", "+x", shell)
	cmd.Run()
	cmd = exec.Command("/bin/bash", "-c", shell)
	cmd.Stdin = os.Stdin
	//捕获标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for true {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

func main() {
	ReadCmdOut()
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
