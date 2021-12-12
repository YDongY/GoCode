package main

import "fmt"

// init 函数可以用来初始化，在 main 函数之前调用，任何一个源文件都可以包含 init 函数
func init() {
	fmt.Println("init1...")
}

func init() {
	fmt.Println("init2...")
}

func main() {
	fmt.Println("main...")
}

/**
对于导包，变量定义，init 函数，main 函数，执行流程:
	包文件变量定义-->包文件 init 函数---> main 文件变量定义---> main 文件 init 函数---> main 文件 main 函数
*/

