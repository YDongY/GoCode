package main

import "fmt"

func main() {

	// 常量必须赋值
	const num int = 100
	// num = 200 // 常量不能修改

	// 常量只能修饰 bool , 数值 , 字符串

	const num2 = num / 2 // 正确 num 必须是常量，得到的结果才是常量，否则编译器报错

	const (
		a    = iota       // 表示 a=0
		b                 // 一行递增一次 b=1
		c, d = iota, iota // c=2 d=2
		e    = iota       // e=3
	)

	fmt.Println(a, b, c, d, e) // 0 1 2 2 3

	// 常量仍然通过字母大小写来判断是否可以跨包
}
