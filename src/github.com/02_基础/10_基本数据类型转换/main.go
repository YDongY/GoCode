package main

import (
	"fmt"
	"strconv"
)

func main() {

	//基本数据类型相互转换,不能自动转换,必须强制转换（显式转换）
	i := 42
	fmt.Printf("%T \n", i) // int
	// 基本数据类型相互转换,不能自动转换,必须强制转换（显式转换）
	// 转换完成之后,i的数据类型并没有改变

	f := float64(i)
	fmt.Printf("f的类型= %T,值=%f\n", f, f) // f的类型= float64,值=42.000000

	var m float32 = float32(f)
	fmt.Printf("m的类型= %T,值=%f\n", m, m) // m的类型= float32,值=42.000000

	i2 := 12.34
	f2 := int64(i2)
	fmt.Printf("f2的类型= %T,值=%d\n", f2, f2) // f2的类型= int64,值=12

	//注意：
	//	 在转换中int64转成int8，编译不会报错，只是转换的结果按溢出处理

	var a int64 = 1234567
	var b = int8(a)
	fmt.Printf("b的类型= %T,值=%d\n", b, b) // b的类型= int8,值=-121

	//-------------------------------------------------------------------------------------

	//基本数据类型与string类型相互转换
	var num1 int = 99
	var num2 float64 = 23.456
	var boo bool = true
	// var char byte = 'h'
	var str string

	// str = fmt.Sprintf("%c", char)
	// fmt.Printf("str type %T str=%q\n", str, str)

	//------------------------------------------------------------------------------------

	//函数转换(基本数据类型转string类型)

	str = strconv.Itoa(num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(boo)
	fmt.Printf("str type %T str=%q\n", str, str)

	//函数转换(string类型转换基本数据类型)

	var str2 string = "true"
	var bo bool
	bo, _ = strconv.ParseBool(str2)
	fmt.Printf("bo type %T b=%v\n", bo, bo)

	var str3 string = "123456"
	var numInt64 int64
	numInt64, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Printf("numInt64 type %T numInt64=%v\n", numInt64, numInt64)

	var str4 string = "123.456"
	var numFloat64 float64
	numFloat64, _ = strconv.ParseFloat(str4, 64)
	fmt.Printf("numFloat64 type %T numFloat64=%v\n", numFloat64, numFloat64)

	//go 对于不能转换成功的，会转换成默认值
	// 例:hello转int类型，结果为0
	// 例:hello转bool类型，结果为false
	var str5 string = "hello"
	// var num int64
	var num int64 = 11
	num, _ = strconv.ParseInt(str5, 10, 64)
	fmt.Printf("num type %T num=%v\n", num, num) //0

}
