package main

import (
	"fmt"
	"strconv"
)

func Parse() {
	// 接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s) // bool, true
	}

	// ------ ParseInt 返回字符串表示的整数值，接受正负号 ------
	// base 指定进制（2到36），如果 base 为0，则会从字符串前置判断，”0x”是16进制，”0”是 8 进制，否则是 10 进制；
	// bitSize 指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64, -354634382
	}
	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64, -3546343826724305832
	}

	// ------ ParseUint 类似ParseInt但不接受正负号，用于无符号整型 ------
	vu32 := "354634382"
	if s, err := strconv.ParseInt(vu32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64, 354634382
	}

	// ------ ParseFloat 返回字符串表示的浮点数 ------
	// s 取值：nan、NaN、inf、+Inf、-Inf、-0 、+0 ......
	f := "3.1415926535"
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415927410125732
	}
}

func Format() {
	/**
	Format 系列函数实现了将给定类型数据格式化为 string 类型数据的功能。
	*/
	v := int64(-42)
	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10) // string, -42
	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16) // string, -2a

	f := 3.1415926535
	s32 := strconv.FormatFloat(f, 'e', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32) // string, 3.1415927e+00
	s64 := strconv.FormatFloat(f, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64) // string, 3.1415926535E+00

	b := true
	s := strconv.FormatBool(b)
	fmt.Printf("%T, %v\n", s, s) // string, true
}

func Append() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b)) // bool:true

	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) // int (base 10):-42
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) // int (base 16):-2a

	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'e', -1, 32)
	fmt.Println(string(b32)) // float32:3.1415927e+00
	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64)) // float64:3.1415926535E+00

	q := []byte("quote:")
	b = strconv.AppendQuote(q, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b)) // quote:"\"Fran & Freddie's Diner\""
}

func main() {

	// --------- Atoi 字符串转换为整数 ---------
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v \n", s, s)
	}

	// --------- Itoa 整数转换字符串 ---------
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v \n", s, s)

	Parse()
	Format()
	Append()
}
