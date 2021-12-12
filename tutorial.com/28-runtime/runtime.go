package main

import (
	"fmt"
	"runtime"
	"time"
)

func TestGosched() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 10; i++ {
		// 切一下，再次分配任务
		if i != 2 {
			runtime.Gosched()
		}
		fmt.Println("hello")
	}
}

func TestGoexit() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
	/**
	输出:
		B.defer
		A.defer
	*/
}

func TestGOMAXPROCS() {
	// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数

	runtime.GOMAXPROCS(1)
	go func() {
		for i := 1; i < 10000; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		for i := 1; i < 10000; i++ {
			fmt.Println("B:", i)
		}
	}()
	time.Sleep(time.Second)
}

func main() {
	// TestGosched()
	// TestGoexit()
	TestGOMAXPROCS()
}
