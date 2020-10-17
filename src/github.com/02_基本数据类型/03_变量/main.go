package main

import "fmt"

func main() {
	var a int = 10
	a = 20
	a = 30

	//修改变量的值，必须为同一数据类型
	var b string = "hello world"
	b = "hello go"

	var c = "go go go..."
	c = "let's go"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//同一作用域不能出现多个相同的变量名，不同类型也不可以
	//var a string = "aaaa"
	//fmt.Println(a)

	e := 100
	fmt.Printf("type e %T,%d", e, e)
}
