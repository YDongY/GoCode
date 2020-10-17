package main

import (
	"fmt"
)

func main() {

	var num int = 10

	//获取变量num的地址
	fmt.Println("num的值=", num)
	fmt.Println("num的地址=", &num)

	fmt.Println("－－－－－－－－－－－－－－－－－－－")
	var ptr *int = &num
	fmt.Println("ptr的地址=", &ptr)
	fmt.Println("ptr的值=", ptr)
	fmt.Println("ptr指向的值=", *ptr)

	fmt.Println("－－－－－－－－－－－－－－－－－－－")
	//修改指针指向的num的值
	*ptr = 20
	fmt.Println("num的值=", num)
	fmt.Println("num的地址=", &num)
	fmt.Println("ptr的地址=", &ptr)
	fmt.Println("ptr的值=", ptr)
	fmt.Println("ptr指向的值=", *ptr)

	fmt.Println("－－－－－－－－－－－－－－－－－－－")

	var num2 int = 100
	ptr = &num2
	fmt.Println("ptr的地址=", &ptr)
	fmt.Println("ptr的值=", ptr)
	fmt.Println("ptr指向的值=", *ptr)
}
