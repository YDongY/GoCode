package main

import "fmt"

func main() {

	//基本数据类型默认值
	var str string //""

	fmt.Println(str)

	//默认int类型
	var a int

	//默认fload64类型
	var b float64
	var c bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//格式化输出
	//整型，原值，浮点型 %d，%v,%f
	fmt.Printf("a=%d,b=%f,b原值=%v,c=%v\n", a, b, b, c)
}
