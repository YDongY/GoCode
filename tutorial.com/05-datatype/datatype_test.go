package _5_datatype

import (
	"math"
	"math/cmplx"
	"testing"
	"unsafe"
)

func TestComplexType(t *testing.T) {
	/** 复数 i = 根号 -1 则 i^2 = -1 i^3 = -i i^4 = 1... */

	var c complex128 = 3 + 4i // 默认 complex128
	t.Log(cmplx.Abs(c))       // 5

	// 欧拉公式 : e^(iπ) + 1 = 0
	t.Log(cmplx.Pow(math.E, 1i*math.Pi) + 1)   // (0+1.2246467991473515e-16i)
	t.Log(cmplx.Exp(1i*math.Pi) + 1)           // (0+1.2246467991473515e-16i)
	t.Logf("%.3f \n", cmplx.Exp(1i*math.Pi)+1) // (0.000+0.000i)
}

func TestIntType(t *testing.T) {
	// ------------- 有符号 ------------
	var (
		i   int
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	)

	t.Logf("%T=%d,%T,%T,%T,%T\n", i, unsafe.Sizeof(i), i8, i16, i32, i64)

	// ----------- 无符号 -----------
	var (
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
	)

	t.Logf("%T=%d,%T,%T,%T,%T\n", ui, unsafe.Sizeof(ui), ui8, ui16, ui32, ui64)
}

func TestFloatType(t *testing.T) {
	t.Logf("minFloat:%f \n", math.MaxFloat32)
	t.Logf("maxFloat:%f \n", math.MaxFloat64)

	// 浮点数据默认 float64 类型
	t.Logf("数据类型: %T \n", 1.1) // float64

	f1 := .123 // 0.123
	f2 := 0.123

	// 科学计数法
	f3 := 5.1234e+2 // 5.1234 * 10^2
	f4 := 5.1234e2  // 5.1234 * 10^2
	f5 := 5.1234e2  // 5.1234 * 10^2
	f6 := 5.1234e+2 // 5.1234 * 10^2
	f7 := 5.1234e-2 // 5.1234 / 10^2
	f8 := 5.1234e-2 // 5.1234 / 10^2

	t.Log(f1, f2, f3, f4, f5, f6, f7, f8)
}

func TestByteType(t *testing.T) {
	var b1 byte = 'a'
	var b2 byte = '0'

	//　输出 ascii 码
	t.Log(b1, b2) // 97 48
	// 输出对应字符，需要使用格式化输出
	t.Logf("b1=%c b2=%c \n", b1, b2) // b1=a b2=0

	// var b3 byte = '北' // overflow 溢出
	var b3 rune = '北'
	t.Log(b3, unsafe.Sizeof(b3)) // 21271 4

	var b4 int = 22269
	t.Logf("b4=%c \n", b4) // 国 输出对应的 unicode 字符

	var b5 = 10 + 'a'
	t.Log(b5) // 107
}

func TestBoolType(t *testing.T) {
	// 只能用 true 和 false ,不可以用 0 或非 0 替代
	var b = false
	t.Log(b, unsafe.Sizeof(b)) // false 1
}

func TestStringType(t *testing.T) {
	// 字符串一旦赋值，不可修改，不可变类型
	t.Log("hello\tworld")            // 双引号，会识别转义字符
	t.Log("hello\\world")            // 多种符号嵌套，需要转义
	t.Log(`hello\nworld`)            // 原生输出，使用反引号(``)
	t.Log("hello" + "world")         // 字符串拼接
	t.Log("hello" + "world" + "go" + // 当字符串，可换行,+ 号必须在结尾
		"python" + "c" + "java" + "php" +
		"vb" + "js" + "c#" + "object-c")

	// string 是不可变 byte slice
	t.Log("\xE4\xB8\xA5" == "严")            // 可以存储任何二进制数据
	t.Log(len("中") == 3, len("hello") == 5) // byte 数

	// Unicode 是一种字符编码，UTF8 是 unicode 的存储实现
	c := []rune("中")
	t.Log(len(c))                  // 字符长度
	t.Logf("中 unicode %x\n", c[0]) // 中 unicode 4e2d
	t.Logf("中 unt8 %x\n", "中")     // 中 unt8 e4b8ad

	/**
	  字符					"中"
	  Unicode				0x4e2d
	  UTF-8				0xe4b8ad
	  string/[]byte		[0xe4 0xb8 0xad]
	*/

	// ----------- 字符串遍历 ------------------
	for _, c := range "中华人民共和国" {
		t.Logf("%[1]d %[1]c \n", c)
	}

	s := []rune("中华人民共和国")
	for i := 0; i < len(s); i++ {
		t.Logf("%c", s[i])
	}
}
