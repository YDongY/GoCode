package utils

import "fmt"

func Say() {
	fmt.Println("hello world")
}

func Reverse(a int, b int) (int, int) {
	return b, a
}
