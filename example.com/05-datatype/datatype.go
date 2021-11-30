package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unsafe"
)

func complexType() {
	/** 复数 i = 根号 -1 则 i^2 = -1 i^3 = -i i^4 = 1... */

	c := 3 + 4i
	fmt.Println(cmplx.Abs(c)) // 5
	// 欧拉公式 : e^(iπ) + 1 = 0
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // (0+1.2246467991473515e-16i)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         // (0+1.2246467991473515e-16i)
	fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1) // (0.000+0.000i)
}

func intType() {
	// ------------- 长度 ------------
	var a int8
	var b int16
	var c int32
	var d int64

	var e int

	fmt.Println(a, b, c, d, e)
	fmt.Println(unsafe.Sizeof(e)) // 8

	// ----------- 符号 -----------
	var a2 uint8
	var b2 uint16
	var c2 uint32
	var d2 uint64

	var e2 uint

	fmt.Println(a2, b2, c2, d2, e2)
	fmt.Println(unsafe.Sizeof(e2))
}

func floatType() {
	fmt.Printf("minFloat:%f \n", math.MaxFloat32)
	fmt.Printf("maxFloat:%f \n", math.MaxFloat64)

	// 浮点数据默认 float64 类型
	var num3 = 1.1
	fmt.Printf("num3 的数据类型%T", num3) // float64

	fmt.Println()
	num4 := .123 //==>0.123
	fmt.Println(num4)

	// 科学计数法
	num5 := 5.1234e+2  //==>5.1234 * 10^2
	num5_2 := 5.1234e2 //==>5.1234 * 10^2
	num5_3 := 5.1234e2 //==>5.1234 * 10^2
	num6 := 5.1234e+2  //==>5.1234 * 10^2
	num7 := 5.1234e-2  //==>5.1234 / 10^2
	num8 := 5.1234e-2  //==>5.1234 / 10^2
	fmt.Println(num5, num5_2, num5_3, num6, num7, num8)
}

func byteType() {
	var c1 byte = 'a'
	var c2 byte = '0'

	//　输出 ascii 码
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	// 输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	// var c3 byte = '北' // overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c\n", c3)
	var c4 rune = '北'
	fmt.Println("c4=", c4, unsafe.Sizeof(c4)) // 21271 4

	var c5 int = 22269        // 国　　utf-8 编码
	fmt.Printf("c5=%c\n", c5) // 输出对应的 unicode 字符

	var n1 = 10 + 'a' //字符运算
	fmt.Println("n1=", n1)
}

func boolType() {
	// 只能用 true 和 false ,不可以用 0 或非 0 替代
	var b = false
	fmt.Println("b=", b)
	fmt.Println("b　占用的空间大小:", unsafe.Sizeof(b))
}

func stringType() {
	// 字符串一旦赋值，不可修改，不可变类型
	// 双引号，会识别转义字符
	b := "hello\nworld"
	fmt.Println("b = ", b)
	fmt.Println("b　占用的空间大小:", unsafe.Sizeof(b))

	// 多种符号嵌套，需要转义
	c := "abc\\ndef"
	fmt.Println("c = ", c)

	d := "aaa\"bbb\"ccc"
	fmt.Println("d = ", d)

	// 原生输出，使用反引号(``)
	e := `123\n456`
	fmt.Println("e = ", e)

	f := `aaa"bbb"ccc`
	fmt.Println("f = ", f)

	// 字符串拼接
	str := "hello" + "world"
	fmt.Println("str = ", str)

	// 当字符串，可换行,+ 号必须在结尾
	str2 := "hello" + "world" + "go" +
		"python" + "c" + "java" + "php" +
		"vb" + "js" + "c#" + "object-c"
	fmt.Println("str2 = ", str2)
}

func main() {
	complexType()
	intType()
	floatType()
	byteType()
	boolType()
}
