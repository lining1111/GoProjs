package main

/*
#cgo CXXFLAGS: -std=c++11
//#include "c/myfun.h"
#include "myfun1.h"
//#cgo CFLAGS: -D__CFUN2_H
//extern int add(int a, int b);
#include <stdio.h>
#include "cfun2.h"
int add2(int a, int b){
	return add(a,b);
}
//#cgo LDFLAGS: -Wl,-rpath="./c" -Lc -lmyfun -lstdc++
*/
import "C"
import (
	"GoProjs/mybuffer"
	"fmt"
	"unsafe"

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

type S11 struct {
	a int32
	b byte
	c int16
}

type Foo struct {
	a int
	b int
}

//https://blog.csdn.net/dielucui7698/article/details/101400578?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~aggregatepage~first_rank_ecpm_v1~rank_v31_ecpm-1-101400578.pc_agg_new_rank&utm_term=cgo+golang+%E4%BC%A0%E9%80%92%E7%BB%93%E6%9E%84%E4%BD%93&spm=1000.2123.3001.4430

func main() {
	//s := S11{a: 10}
	//
	////length := int(unsafe.Sizeof(s))
	////fmt.Println(length)
	//
	//C.CGO_setStruct1((*C.CGO_S11)(unsafe.Pointer(&s)))

	foo := Foo{10, 20}
	foos := []*Foo{&Foo{1, 2}, &Foo{3, 4}}

	C.pass_struct((*C.Foo)(unsafe.Pointer(&foo)))
	//C.pass_structs((**C.Foo)(unsafe.Pointer(&foos[0])), C.int(len(foos)))

	for _, f := range foos {
		fmt.Printf("%d,%d\n", f.a, f.b)
	}

	//
	//s11 := myfun1.S{}
	//myfun1.TestMyFun1(&s11)

	//sum := C.add(1, 2)
	//fmt.Println(sum)
	//
	str := "hello world"
	length1 := C.lenStr(C.CString(str))
	fmt.Println(length1)

	buf := mybuffer.NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))

	buf.Push(15)

	fmt.Println(buf.Pop())

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
