package main

import "fmt"

func main() {
	// 将 bigint 定义为 int64 类型
	type bigint int64

	// 将 int 取一个别名 IntAlias
	type IntAlias = int

	// 将 a 声明为 bigint 类型
	var a bigint
	// 查看 a 的类型名
	fmt.Printf("a type %T \n",a) // a type main.bigint

	// 将 b 声明为 IntAlias 类型
	var b IntAlias
	// 查看 b 的类型名
	fmt.Printf("b type %T \n",b) // int
}
