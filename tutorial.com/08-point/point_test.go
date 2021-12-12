package _8_point

import (
	"testing"
)

func TestPoint(t *testing.T) {
	var p *int
	t.Log(p) // nil
	t.Logf("%p \n", p) // 0x0

	// *p = 123 // panic: runtime error: invalid memory address or nil pointer dereference
}

func TestPoint2(t *testing.T){
	var i = 10
	t.Log("i的地址＝", &i)

	// 指针变量存的是一个地址，这个地址指向的空间存的才是值
	// 1.ptr 是一个指针变量
	// 2.ptr 的类型是 *int
	// 3.ptr 本身的值 &i
	// 4.ptr 本身的地址 &ptr
	var ptr *int = &i
	t.Log("ptr 本身的值 =", ptr)  // 0xc000098010
	t.Log("ptr 的地址 =", &ptr)  // 0xc00000e030
	t.Log("ptr 指向的值 =", *ptr) // 10
}
