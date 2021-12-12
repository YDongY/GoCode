package main // 特殊的包，用来定义一个独立的可执行程序

import (
	"fmt"
	"os"
) // 导包，必须在 package 之后

// main 是程序执行的入口
func main() { // { 必须和关键 func 同一行
	fmt.Println("hello 世界") // 原生支持 unicode

	// 获取命令行参数 os.Args
	if len(os.Args) > 1 {
		fmt.Println("Args:", os.Args[1:])
	}

	// main 函数不支持返回值，通过 os.Exit 获取返回状态
	os.Exit(0)
}
