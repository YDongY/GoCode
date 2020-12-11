package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	map1 = make(map[int]int, 10)
	// 全局互斥锁
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	map1[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 15; i++ {
		go test(i)
	}
	time.Sleep(time.Second * 5)
	// 遍历结果
	lock.Lock()
	for _, v := range map1 {
		fmt.Println(v)
	}
	lock.Unlock()
}
