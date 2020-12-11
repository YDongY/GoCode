package main

import "fmt"

func dummy(b int) int{
	var c int
	b = c
	return c
}

func void(){

}

func main() {
	var a int
	void()
	fmt.Println(a,dummy(0))

	// 命令执行：go run -gcflags "-m -l" main.go
}
