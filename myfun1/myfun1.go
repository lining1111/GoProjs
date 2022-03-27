/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: myfun1.i

package myfun1

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_myfun1_b7e9c6cff453cebb(swig_intgo arg1);
extern void _wrap_S11_a_set_myfun1_b7e9c6cff453cebb(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_S11_a_get_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_S11_b_set_myfun1_b7e9c6cff453cebb(uintptr_t arg1, char arg2);
extern char _wrap_S11_b_get_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_S11_c_set_myfun1_b7e9c6cff453cebb(uintptr_t arg1, short arg2);
extern short _wrap_S11_c_get_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern uintptr_t _wrap_new_S11_myfun1_b7e9c6cff453cebb(void);
extern void _wrap_delete_S11_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_CGO_S11_a_set_myfun1_b7e9c6cff453cebb(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_CGO_S11_a_get_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_CGO_S11_b_set_myfun1_b7e9c6cff453cebb(uintptr_t arg1, char arg2);
extern char _wrap_CGO_S11_b_get_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_CGO_S11_c_set_myfun1_b7e9c6cff453cebb(uintptr_t arg1, short arg2);
extern short _wrap_CGO_S11_c_get_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern uintptr_t _wrap_new_CGO_S11_myfun1_b7e9c6cff453cebb(void);
extern void _wrap_delete_CGO_S11_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_setStructS11_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
extern void _wrap_setStructCGO_S11_myfun1_b7e9c6cff453cebb(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_myfun1_b7e9c6cff453cebb(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrS11 uintptr

func (p SwigcptrS11) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrS11) SwigIsS11() {
}

func (arg1 SwigcptrS11) SetA(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_S11_a_set_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrS11) GetA() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_S11_a_get_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrS11) SetB(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_S11_b_set_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrS11) GetB() (_swig_ret byte) {
	var swig_r byte
	_swig_i_0 := arg1
	swig_r = (byte)(C._wrap_S11_b_get_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrS11) SetC(arg2 int16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_S11_c_set_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0), C.short(_swig_i_1))
}

func (arg1 SwigcptrS11) GetC() (_swig_ret int16) {
	var swig_r int16
	_swig_i_0 := arg1
	swig_r = (int16)(C._wrap_S11_c_get_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewS11() (_swig_ret S11) {
	var swig_r S11
	swig_r = (S11)(SwigcptrS11(C._wrap_new_S11_myfun1_b7e9c6cff453cebb()))
	return swig_r
}

func DeleteS11(arg1 S11) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_S11_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0))
}

type S11 interface {
	Swigcptr() uintptr
	SwigIsS11()
	SetA(arg2 int)
	GetA() (_swig_ret int)
	SetB(arg2 byte)
	GetB() (_swig_ret byte)
	SetC(arg2 int16)
	GetC() (_swig_ret int16)
}

type SwigcptrCGO_S11 uintptr

func (p SwigcptrCGO_S11) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrCGO_S11) SwigIsCGO_S11() {
}

func (arg1 SwigcptrCGO_S11) SetA(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CGO_S11_a_set_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrCGO_S11) GetA() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_CGO_S11_a_get_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCGO_S11) SetB(arg2 byte) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CGO_S11_b_set_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0), C.char(_swig_i_1))
}

func (arg1 SwigcptrCGO_S11) GetB() (_swig_ret byte) {
	var swig_r byte
	_swig_i_0 := arg1
	swig_r = (byte)(C._wrap_CGO_S11_b_get_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrCGO_S11) SetC(arg2 int16) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_CGO_S11_c_set_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0), C.short(_swig_i_1))
}

func (arg1 SwigcptrCGO_S11) GetC() (_swig_ret int16) {
	var swig_r int16
	_swig_i_0 := arg1
	swig_r = (int16)(C._wrap_CGO_S11_c_get_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func NewCGO_S11() (_swig_ret CGO_S11) {
	var swig_r CGO_S11
	swig_r = (CGO_S11)(SwigcptrCGO_S11(C._wrap_new_CGO_S11_myfun1_b7e9c6cff453cebb()))
	return swig_r
}

func DeleteCGO_S11(arg1 CGO_S11) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_CGO_S11_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0))
}

type CGO_S11 interface {
	Swigcptr() uintptr
	SwigIsCGO_S11()
	SetA(arg2 int)
	GetA() (_swig_ret int)
	SetB(arg2 byte)
	GetB() (_swig_ret byte)
	SetC(arg2 int16)
	GetC() (_swig_ret int16)
}

func SetStructS11(arg1 S11) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_setStructS11_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0))
}

func SetStructCGO_S11(arg1 CGO_S11) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_setStructCGO_S11_myfun1_b7e9c6cff453cebb(C.uintptr_t(_swig_i_0))
}
