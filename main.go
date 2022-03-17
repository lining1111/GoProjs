package main

/*
//#include "c/myfun.h"
//#include "c/myfun1.h"
//#cgo CFLAGS: -D__CFUN2_H
//extern int add(int a, int b);
#include "cfun2.h"
int add2(int a, int b){
	return add(a,b);
}
//#cgo LDFLAGS: -Wl,-rpath="./c" -Lc -lmyfun -lstdc++
*/
import "C"
import (
	"fmt"
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

type S1 struct {
	a int
}

func setS1(s1 *S1) {
	s1.a = 10
}

func main() {
	//var s C.S1
	//s.a = 5
	//length := unsafe.Sizeof(s)
	//fmt.Println(length)

	//C.setStruct(&s)
	//
	//s11 := myfun1.S{}
	//myfun1.TestMyFun1(&s11)

	sum := C.add(1, 2)
	fmt.Println(sum)

	str := "hello world"
	length := C.lenStr(C.CString(str))
	fmt.Println(length)

	//s1 := S1{
	//}
	//setS1(&s1)
	//
	////learn input
	////learn.InputScan()
	////learn img
	//learn.Image()

	//http.Handle("/",http.FileServer(http.Dir("html")))
	//http.HandleFunc("/hello",hello)
	//http.HandleFunc("/process",process)
	//
	//defer http.ListenAndServe("localhost:8080",nil)

	//t,_:=template.ParseFiles("html/tmpl.html")

	//person := test.Person{"lining",34}
	//fmt.Println("name ",person.Name, " age ",person.Age)
	//
	//test.TestDecXml("xml/servers.xml")
	//test.TestEncXml()
	//
	//test.TestTmpl()

	//slice := make([]int, 4, 10)
	//for i := 0; i < len(slice); i++ {
	//	slice[i] = i + 1
	//}
	//fmt.Println("slice = ", slice)
	//config := test.Config{
	//	Host:    "127.0.0.1",
	//	Port:    3306,
	//	Usr:     "root",
	//	Passwd:  "admin",
	//	DbName:  "parking_v3",
	//	Charset: "utf8",
	//}
	//test.Open(config)
	//test.GetDeviceInfo()

}
