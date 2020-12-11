package main

import (
	"fmt"
	"sync/atomic"
)

var (
	seq int64
)

func GenID() int64 {
	return atomic.AddInt64(&seq, 1)
}

func main() {
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())

	// --race 参数，开启运行时（runtime）对竞态问题的分析
	// go run --race main.go
}
