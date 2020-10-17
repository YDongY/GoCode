package main

import "fmt"

// 1,1,2,3,5,8,13
func fib(n int) int {
	if n <= 2 {
		return 1
	} else {
		return fib(n-2) + fib(n-1)
	}
}

func main() {

	res := fib(7)
	fmt.Println(res)

}
