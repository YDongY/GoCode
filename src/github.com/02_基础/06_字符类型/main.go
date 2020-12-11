package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	//　输出ascii码
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	//输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	// var c3 byte = '北' //overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c\n", c3)
	var c4 rune = '北'
	fmt.Println("c4=", c4, unsafe.Sizeof(c4)) // 21271 4

	var c5 int = 22269        //===>国　　utf-8编码
	fmt.Printf("c5=%c\n", c5) //输出对应的unicode字符

	var n1 = 10 + 'a' //字符运算
	fmt.Println("n1=", n1)
}
