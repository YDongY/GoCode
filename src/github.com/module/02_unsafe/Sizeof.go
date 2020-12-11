package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//查看数据所占字节数和数据类型
	var f = "aaa"
	var g = 222
	fmt.Printf("f 的类型是%T 占用的字节数%d ", f, unsafe.Sizeof(f))
	fmt.Printf("\n")
	fmt.Printf("g 的类型是%T 占用的字节数%d ", g, unsafe.Sizeof(g))
}
