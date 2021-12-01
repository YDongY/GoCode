package main

import (
	"fmt"
	"sync"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func add() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func Mutex() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func RWMutex() {
	// 读写锁分为两种：读锁和写锁。
	// 当一个 goroutine 获取读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
	// 当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。
}

func main() {
	Mutex()
}
