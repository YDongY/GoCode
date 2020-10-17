package main

import "fmt"

var (
	Fun1 = func(n1 int, n2 int) int { // 全局匿名函数
		return n1 + n2
	}
)

func main() {
	// 使用匿名函数，求和
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println(res)

	// 匿名函数赋值给变量
	res2 := func(n1 int, n2 int) int {
		return n1 + n2
	}

	fmt.Println(res2(10, 20))

	fmt.Println(Fun1(100, 200))
}
