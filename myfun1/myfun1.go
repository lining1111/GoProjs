package myfun1

/*
#include "../c/myfun1.h"
#cgo LDFLAGS: -Wl,-rpath="../c" -L../c -lmyfun1
*/
import "C"

type S struct {
	a int
	b byte
}

func TestMyFun1(s *S)  {
	s1 :=C.S11{}
	C.setStruct1(&s1)
	s.a = int(s1.a)
	s.b = byte(s1.b)
}