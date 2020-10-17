package main

import "fmt"

func addUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

func main() {
	fmt.Println(addUpper()(2))
	s := addUpper()

	fmt.Println(s(1)) // 11
	fmt.Println(s(2)) // 13

}
