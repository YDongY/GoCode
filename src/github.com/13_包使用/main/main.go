package main

import (
	"fmt"

	// "github.com/13_包使用/conf"
	cf "github.com/13_包使用/conf" // 引包起别名，需要用别名访问
	"github.com/13_包使用/utils"
)

func main() {
	utils.Say()
	fmt.Println(utils.Reverse(10, 20))
	fmt.Println(cf.Num) // 使用别名访问
}

// 注意同一个包下不能存在相同的声明
// main 包 只能有一个