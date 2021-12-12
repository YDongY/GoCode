package _6_datatype_conversion

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIntToFloat(t *testing.T) {
	// 基本数据类型相互转换,不能自动转换,必须强制转换（显式转换）
	var i = 100
	var f float64 = float64(i)
	t.Logf("i=(%d,%T),f=(%f,%T)\n", i, i, f, f) // i=(100,int),f=(100.000000,float64)
}

func TestFloatToInt(t *testing.T) {
	var f = 3.14
	var i = int(f)
	t.Logf("i=(%d,%T),f=(%f,%T)\n", i, i, f, f) // i=(3,int),f=(3.140000,float64)
}

func TestInt64ToInt8(t *testing.T) {
	// 在转换中 int64 转成 int8，编译不会报错，只是转换的结果按溢出处理
	var i64 int64 = 1234567
	var i8 = int8(i64)
	t.Logf("i64=(%d,%T),i8=(%d,%T)\n", i64, i64, i8, i8) // i64=(1234567,int64),i8=(-121,int8)
}

func TestToString(t *testing.T) {
	//基本数据类型与 string 类型相互转换
	var i = 100
	var f = 3.14
	var b = true
	var c = 'a'

	t.Log(fmt.Sprintf("%d", i), fmt.Sprintf("%f", f), fmt.Sprintf("%v", b), fmt.Sprintf("%c", c))
}

func TestToStringByFunc(t *testing.T) {
	// 函数转换(基本数据类型转 string 类型)

	var str string
	str = strconv.Itoa(100)
	t.Logf("str type %T str=%q\n", str, str)

	str = strconv.FormatInt(int64(100), 10)
	t.Logf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(3.14, 'f', 10, 64)
	t.Logf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(true)
	t.Logf("str type %T str=%q\n", str, str)
}

func TestStringToByFunc(t *testing.T) {
	// 函数转换(string 类型转换基本数据类型)

	var b bool
	b, _ = strconv.ParseBool("true")
	t.Logf("b type %T b=%v\n", b, b)

	var i64 int64
	i64, _ = strconv.ParseInt("123456", 10, 64)
	t.Logf("i64 type %T i64=%v\n", i64, i64)

	var f64 float64
	f64, _ = strconv.ParseFloat("123.456", 64)
	t.Logf("f64 type %T f64=%v\n", f64, f64)
}

func TestError(t *testing.T) {
	// go 对于不能转换成功的，会转换成零值
	i64, err := strconv.ParseInt("hello", 10, 64)
	if err != nil {
		t.Log("ParseInt error:", err) // ParseInt error: strconv.ParseInt: parsing "hello": invalid syntax
	}
	t.Logf("i64 type %T i64=%v\n", i64, i64) // i64 type int64 i64=0
}
