package main

import (
	"fmt"
	_ "unsafe" //如果没有使用一个包，可以加一个_，表示忽略
)

func main() {
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1 + 20)
	n3 = int8(n1 + 20)

	fmt.Printf("n2=%d,n3=%d\n", n2, n3)

	var a int32 = 12
	var b int8
	var c int8
	b = int8(a) + 127
	// c = int8(a) + 128 //编译不通过
	fmt.Printf("b=%d\n", b)
	fmt.Printf("c=%d\n", c)
}
