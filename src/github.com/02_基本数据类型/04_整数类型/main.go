package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//============重点============
	//*********有符号*********
	// int8: -128~127
	// int16,int32,int64
	var a int8 = 127
	fmt.Println(a)

	//******无符号******
	//uint8: 0-255
	//uint16,uint32,uint64
	var b uint8 = 255
	fmt.Println(b)

	//其他整型
	// int,uint,rune,byte
	//其中int,uint与操作系统有关

	var c int = 1000000
	fmt.Println(c)

	//byte: 0-255
	var d byte = 255
	fmt.Println(d)

	var e rune = 123456
	fmt.Println(e)

	//%%%%%%%%%%%%%%%%%查看变量数据类型%%%%%%%%%%%%%%%%%%%%%
	var f = 1000
	fmt.Printf("f 的类型是%T 占用的字节数%d", f, unsafe.Sizeof(f))
}
