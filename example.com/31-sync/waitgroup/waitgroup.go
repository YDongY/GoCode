package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func add() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
}