package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("printf格式化输出\n")
	fmt.Println("随机整数:", rand.Intn(10))
	fmt.Println("开方", math.Sqrt(4))
}
