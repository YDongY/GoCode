package main

import "fmt"

//声明全局变量
var (
	_a = 10
	_b = 20
	_c = 30
)

func main() {
	//定义变量
	var name string
	//变量赋值
	name = "ydy"
	//使用变量
	fmt.Println("name=", name)

	//省略var ,等加于 var name2 string&name2="tom"
	name2 := "tom"

	fmt.Println("name2=", name2)

	//一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//var a1, n, a3 = 100, "tom", 888
	//fmt.Println("a1=", a1, "n=", n, "a3=", a3)

	a1, n, a3 := 100, "tom", 888
	fmt.Println("a1=", a1, "n=", n, "a3=", a3)

	//使用全局变量
	fmt.Println("_a=", _a, "_b=", _b, "_c=", _c)
}
