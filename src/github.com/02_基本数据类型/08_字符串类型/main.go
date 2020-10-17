package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//字符串一旦赋值，不可修改，不可变类型
	//双引号，会识别转义字符
	b := "hello\nworld"
	fmt.Println("b = ", b)
	fmt.Println("b　占用的空间大小:", unsafe.Sizeof(b))

	//多种符号嵌套，需要转义
	c := "abc\\ndef"
	fmt.Println("c = ", c)

	d := "aaa\"bbb\"ccc"
	fmt.Println("d = ", d)

	//原生输出，使用反引号(``)
	e := `123\n456`
	fmt.Println("e = ", e)

	f := `aaa"bbb"ccc`
	fmt.Println("f = ", f)

}
