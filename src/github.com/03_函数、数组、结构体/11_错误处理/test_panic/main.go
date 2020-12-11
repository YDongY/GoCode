package main

import "fmt"

func main() {

	// panic 是 Go 语言中的一个 内置函 数，它终止正常的控制流程并引发恐慌（ panicking ) , 导致程序停止执行
	fmt.Println("---main---")
	panic("main panic")
	fmt.Println("panic after")
}
