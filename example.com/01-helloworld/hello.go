package main // 特殊的包，用来定义一个独立的可执行程序

import "fmt" // 导包，必须在 package 之后

// main 是程序执行的入口
func main() { // { 必须和 关键在 func 同一行
	fmt.Println("hello 世界") // 原声支持 unicode
	// Println 输出
}
