package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// ------------- 长度 ------------
	var a int8
	var b int16
	var c int32
	var d int64

	var e int

	fmt.Println(a,b,c,d,e)
	fmt.Println(unsafe.Sizeof(e)) // 8

	// ----------- 符号 -----------
	var a2 uint8
	var b2 uint16
	var c2 uint32
	var d2 uint64

	var e2 uint

	fmt.Println(a2,b2,c2,d2,e2)
	fmt.Println(unsafe.Sizeof(e2))
}
