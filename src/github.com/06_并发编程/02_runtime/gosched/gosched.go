package main

import (
	"fmt"
	"runtime"
)

func main() {

	// 创建一个goroutine
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched() // 让出 CPU 执行权
		fmt.Println("hello")
	}

}
