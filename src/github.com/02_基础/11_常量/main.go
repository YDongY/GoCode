package main

import "fmt"

func testIota1() {

	const (
		a    = iota       // 表示 a=0
		b                 // 一行递增一次 b=1
		c, d = iota, iota // c=2 d=2
		e    = iota       // e=3
	)

	fmt.Println(a, b, c, d, e) // 0 1 2 2 3

}

func testIota2() {
	const (
		n1 = iota // 0
		n2        // 1
		_
		n4 // 3
	)
	fmt.Println(n1, n2, n4) // 0 1 3
}

func testIota3() {
	const (
		n1 = iota // 0
		n2 = 100  // 100
		n3 = iota // 2
		n4        // 3
	)
	const n5 = iota                 //0
	fmt.Println(n1, n2, n3, n4, n5) // 0 100 2 3 0
}

func testIota4() {
	const (
		i = iota
		a = 1 << iota
		b
		c
	)
	fmt.Println(i, a, b, c) // 0 2 4 8
}

func testIota5() {
	const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	fmt.Println(a, b, c, d, e, f) // 1 2 2 3 3 4
}

func main() {
	/**
	- 常量必须赋值
	- 常量只能修饰 bool , 数值 , 字符串
	- 常量仍然通过字母大小写来判断是否可以跨包
	*/

	const num int = 100
	// num = 200 // 常量不能修改
	const num2 = num / 2 // 正确 num 必须是常量，得到的结果才是常量，否则编译器报错

	const (
		n1 = 100
		n2
		n3
	)

	fmt.Println(n1, n2, n3) // 100 100 100

	testIota1()
	testIota2()
	testIota3()
	testIota4()
	testIota5()
}
