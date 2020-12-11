package main

import "fmt"

func main() {
	var a int8 = 3
	var b int8 = -2

	// 左移
	fmt.Println(a << 1) // 6
	fmt.Println(b << 1) // -4

	// 右移
	fmt.Println(a >> 1) // 1
	// 负数右移最小值为 -1，操作原码
	fmt.Println(b >> 1) // -1

	// 负数按照补码计算: | & ^

	fmt.Println(a | b) // -1

	fmt.Println(a & b) // 2

	fmt.Println(a ^ b) // -3

	fmt.Println(a &^ b) // 1
}
