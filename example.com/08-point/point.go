package main

import (
	"fmt"
)

func testPoint() {
	var p *int
	fmt.Println(p) // nil

	fmt.Printf("%p", p) // 0x0

	// *p = 123 // panic: runtime error: invalid memory address or nil pointer dereference
}

func main() {
	var i int = 10
	// i 的地址
	fmt.Println("i的地址＝", &i) //0xc000098010

	// 指针变量存的是一个地址，这个地址指向的空间存的才是值
	// 1.ptr 是一个指针变量
	// 2.ptr 的类型是 *int
	// 3.ptr 本身的值 &i
	// 4.ptr 本身的地址 &ptr
	var ptr *int = &i
	fmt.Println("ptr本身的值=", ptr)  // 0xc000098010
	fmt.Println("ptr的地址=", &ptr)  // 0xc00000e030
	fmt.Println("ptr指向的值=", *ptr) // 10

	testPoint()

}
