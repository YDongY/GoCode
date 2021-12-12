package main

import "fmt"

func addUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

func playerGen(name string) func() (string, int) {
	hp := 150
	// 返回创建的闭包
	return func() (string, int) {
		// 将变量引用到闭包中
		return name, hp
	}
}

func main() {
	s := addUpper()

	fmt.Println(s(1)) // 11
	fmt.Println(s(2)) // 13

	s2 := addUpper()   // 重新调用 addUpper() 会重新声明及赋值局部变量 n
	fmt.Println(s2(1)) // 11
	fmt.Println(s2(2)) // 13

	// 闭包生成器
	generator := playerGen("jack")
	name, hp := generator()

	fmt.Println(name, hp)
}
