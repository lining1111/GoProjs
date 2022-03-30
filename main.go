package main

import (
	"GoProjs/myfun1"
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
	a int
	b int
	c int
}

type Foo struct {
	a int
	b byte
}

//https://blog.csdn.net/dielucui7698/article/details/101400578?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~aggregatepage~first_rank_ecpm_v1~rank_v31_ecpm-1-101400578.pc_agg_new_rank&utm_term=cgo+golang+%E4%BC%A0%E9%80%92%E7%BB%93%E6%9E%84%E4%BD%93&spm=1000.2123.3001.4430

func main() {

	//exampleClass.ExampleClassTest()

	//bs := make([]byte, 51)
	//bsp1 := &bs[0]
	//fmt.Printf("%T", bsp1)
	//bsp := (*byte)(unsafe.Pointer(&bs[0]))
	//bspUintptr := uintptr(unsafe.Pointer(bsp)) // 第11行，行号状态机描述的有问题
	//
	//bs[1] = 200
	//fmt.Printf("%d\n", *(*byte)(unsafe.Pointer(bspUintptr + 1)))

	//×byte-->[]byte
	var arr1 = [4]byte{11, 22, 33, 44}

	p := &arr1[0]
	fmt.Printf("%T,%v\n", p, &arr1) // 数组指针
	//先变成任意指针
	ptr := uintptr(unsafe.Pointer(p))
	var arr2 = make([]byte, 4)
	for i := 0; i < len(arr2); i++ {
		//然后和c取值一样
		arr2[i] = *(*byte)(unsafe.Pointer(ptr + uintptr(i)))
	}
	fmt.Println(arr2)

	s11 := myfun1.NewS11()
	myfun1.SetStructS11(s11)

	fmt.Printf("a:%d,b:%d,c:%d\n", s11.GetA(), s11.GetB(), s11.GetC())

	serialBuf := make([]byte, 1024)
	serialBudLen := uint(0)

	serialPtr := &serialBuf[0]
	serialLen := serialBudLen
	myfun1.SerialStructS11Ptr(serialPtr, &serialLen, s11.Swigcptr())
	fmt.Printf("a:%d,b:%d,c:%d\n", s11.GetA(), s11.GetB(), s11.GetC())
	result := serialBuf[0:serialBudLen]

	fmt.Println(result)

	myfun1.DeleteS11(s11)

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
