package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("---哼---")
		time.Sleep(100 * time.Millisecond)
	}
}

func reply() {
	for i := 0; i < 5; i++ {
		fmt.Println("---哈---")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sing()
	go reply()

	for {

	}
}
