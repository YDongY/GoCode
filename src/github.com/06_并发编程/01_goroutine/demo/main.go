package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("---slowFunc---")
}

func main() {
	go slowFunc()
	fmt.Println("---main---")
	time.Sleep(time.Second * 3) // 等待 Goroutine 执行完毕，否则 Goroutine 无法执行就退出了
}
