package main

import "fmt"

func main() {
	// 1.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2.
	var i = 10
	for i > 0 {
		fmt.Println(i)
		i--
	}

	// 3.
	var j = 10
	for {
		fmt.Println(j)
		j--
		if j == 0 {
			break
		}
	}

	// 4.
	var slice = []int{1, 2, 3, 4}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}
