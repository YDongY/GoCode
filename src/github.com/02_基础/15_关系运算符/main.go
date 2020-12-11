package main

import "fmt"

func main() {
	fmt.Println(1 == 1)  // true
	fmt.Println(1 != -1) // true
	fmt.Println(1 < -1)  // false
	fmt.Println(1 > -1)  // true
	fmt.Println(1 >= -1) // true
	fmt.Println(1 <= -1) // false

	flag := 1 > 1
	fmt.Println(flag) // false

	flag2 := 1 == 1
	fmt.Println(flag2) // true

}
