package main

import "fmt"

func main() {
	fmt.Println(1 == 1 && 2 > 1) // true
	fmt.Println(1 == 1 && 2 > 3) // false
	fmt.Println(1 == 1 || 2 > 3) // true
	fmt.Println(1 != 1 || 2 > 3) // false

	fmt.Println(!true)     // false
	fmt.Println(1 != 1)    // false
	fmt.Println(!(1 != 1)) // true

	// 如果 && 左边为假，则不判断右边
	// 如果 || 右边为真，则不判断右边

}
