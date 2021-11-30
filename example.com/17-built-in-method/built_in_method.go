package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("%T,%v,%v \n", num1, num1, &num1) //int,100,0xc0000b8010

	// new 分配值内存
	num2 := new(int)

	fmt.Printf("%T,%v,%v \n", num2, num2, &num2) // *int,0xc0000160e8,0xc00000e030
	fmt.Println(*num2)                           // 0

	// make 分配引用内存
}
