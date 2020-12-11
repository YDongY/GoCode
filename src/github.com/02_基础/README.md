# 入门

## 输入输出

- 输出

```go
func main() {
    // Fprintxx()
    fmt.Fprint(os.Stdout,"内容1") // 向流中写入内容
    fmt.Fprintln(os.Stdout,"内容2") // 向流中写入内容,并加上换行符
    fmt.Fprintf(os.Stdout,"%s","内容3") // 格式化向流中写入内容 
    
    // Printxx()
    fmt.Print("1","2") // 输出内容不换行
    fmt.Println("1","2") // 输出内容并换行符
    fmt.Printf("%s-%s","1","2") // 格式化输出
    
    // Sprintxx()
    ret :=fmt.Sprint("1","2")
    ret :=fmt.Sprintln("1","2")
    ret :=fmt.Sprintf("%s-%s","1","2")
}
```

- 输入

```go
func main() {
	
    // Scanln(&name, &age)
    var name, age string
    fmt.Printf("请输入姓名和年龄:")
    fmt.Scanln(&name, &age)
    fmt.Printf("name = %s age = %s", name, age)

    // Scanf("%s\n%s", &name, &age)
    fmt.Printf("请输入姓名和年龄:")
    fmt.Scanf("%s\n%s", &name, &age)
    fmt.Printf("name = %s age = %s", name, age)
}
```

Sprintxx 、Printxx 、Fprintxx 区别：

- Sprintxx()：把结果以字符串返回
- Printxx()：把结果打印控制台
- Fprintxx()：把结果写入流（文件）

## 转义

```go
package main

import "fmt"

func main() {
	fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	fmt.Printf("%d %x %X %o %d \n", 18, 11, 12, 18, 18) // 18 b C 22 18
	fmt.Printf("%f \n", 1.5)                            // 1.500000
	fmt.Printf("%t \n", false)
	fmt.Printf("%c \n", 97)            // a
	fmt.Printf("%s \n", "hello world") // hello world
	fmt.Printf("%q \n", "hello world") // "hello world"
	fmt.Printf("%v \n", "你好")          // 你好
	fmt.Printf("%T \n", false)         // bool

	i := 5
	fmt.Printf("%p \n", &i)   // 0xc00001a088
	fmt.Printf("%d%% \n", 20) // 20%
}
```

## 变量声明与定义

名称开头字母或下划线，后面是任意数量字符、数字。

- 一个实体名称声明在函数中，则在函数局部有效。
- 声明在函数外，对整个包有效
- 如果声明的实体名称以**大写字母开头**，则表示其可跨包使用，例如：`fmt.Println()`，函数首字母大写

GO 语言编程风格，作用域越小，变量名越短，否则越长，且变量名采用**驼峰式**风格

```go
var name string = "hello" // 声明+初始化

// 或者
var name string // 声明
name = "hello"  // 初始化

// 或者

name := "hello" // 初始化+自动类型检查
```

对于 GO 语言而言定义之后的变量，必须使用，如果不则会编译失败。对于不同类型的变量声明具有默认值。

对于数值类型默认值是 0 、布尔类型默认值为 flase、对于字符串类型默认值为 ""、对于接口和引用类型（slice，指针，map，channel，函数）默认值为 nil、对于像数组或结构体这样的复合类型，默认值是其所有元素的默认值。

### 一次性声明多个变量

```go
var i, j, k int                  // int,int,int
var b, f, s = true, 2.3, "hello" // bool,float64,string
a, b, c := 100, "tom", 888
```

### 变量的修改

```go
// 修改变量的值，必须为同一数据类型
var b string = "hello world"
b = "hello go"

// b :=100 // 同一作用域不能出现多个相同的变量名，不同类型也不可以
```

## 常量和枚举

通过 const 关键字定义常量，声明的常量必须要初始化值，而且常量的值无法修改。如果是通过计算的到的常量，则参与计算的也必须是常量。

```go
// 常量必须赋值
const num int = 100

// num = 200 // 常量不能修改

const num2 = num / 2 // 正确 num 必须是常量，得到的结果才是常量，否则编译器报错
```

```go
const (
    a    = iota       // 表示 a=0
    b                 // 一行递增一次 b=1
    c, d = iota, iota // c=2 d=2
    e    = iota       // e=3
)

fmt.Println(a, b, c, d, e) // 0 1 2 2 3
```

## 复数类型

```go
c := 3 + 4i
fmt.Println(cmplx.Abs(c)) // 5

// 欧拉公式 : e^(iπ) + 1 = 0
fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1) // (0+1.2246467991473515e-16i)
fmt.Println(cmplx.Exp(1i*math.Pi) + 1)         // (0+1.2246467991473515e-16i)
fmt.Printf("%.3f \n", cmplx.Exp(1i*math.Pi)+1) // (0.000+0.000i)
```

## 数据类型转换

基本数据类型相互转换,不能自动转换,必须强制转换（显式转换），格式 `type(变量名)`，而且一般都是类型小的转换类型大的，反之可能会导致结果范围溢出。

对于转换失败的类型，默认会转换为对应类型的默认值。

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 42
	f := float64(i)　// int -> float64 转换完成之后,i 的数据类型并没有改变

	var a int64 = 1234567
	var b int8 = int8(a)
	fmt.Println("b=", b) //-121 // int64 -> int8 范围溢出


	// 基本数据类型与string类型相互转换
	var num1 int = 99
	var num2 float64 = 23.456
	var boo bool = true
	var char byte = 'h'
	var str string

	str = fmt.Sprintf("%c", char) // byte -> string
	fmt.Printf("str type %T str=%q\n", str, str)


	// 基本数据类型转 string 类型(函数转换)

	str = strconv.Itoa(num1)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 10, 64) //
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(boo)
	fmt.Printf("str type %T str=%q\n", str, str)

	// string 类型转换基本数据类型(函数转换

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

	// go 对于不能转换成功的，会转换成默认值
	// 例:hello 转 int 类型，结果为 0
	// 例:hello 转 bool 类型，结果为 false
	var str5 string = "hello"
	// var num int64
	var num int64 = 11
	num, _ = strconv.ParseInt(str5, 10, 64)
	fmt.Printf("num type %T num=%v\n", num, num) //0

}
```

## 指针

变量是存储值的地方，而指针的值是一个变量的地址，一个指针指示值所保存的位置

```go
var i int = 10 // 定义一个 int 类型变量 i
fmt.Println("i的地址＝", &i) // 0xc000098010
j = &i
*j = 2
fmt.Println(i) // 2
```

可以通过 `&name` 来获取变量值所存储的地址，还可以通过 `*name=value`　来修改 `&name` 地址的值。

## 流程控制

### `if-else`

```go
if xxx >= 18 {
	// ...
} else {
	// ...
}
```

```go
if a := 10 * 2; a == 10 {
    fmt.Println("a>10")
} else {
    fmt.Println("a<10")
}
```

### `if-else if-else`

```go
var num int = 500

if num >= 100 && num < 1000 {
	// ...
} else if num >= 1000 {
	// ...
} else {
	// ...
```

### `switch` 每个 case 不用添加 break 自动在结尾添加 break

```go
var key byte
fmt.Println("请输入一个字符：a,b,c,d,e,f,g")
fmt.Scanf("%c", &key)

switch key { // key 常量、变量、返回值
	// case 类型必须与 switch 一致，case 表达式为常量，不能重复
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	case 'd':
		fmt.Println("周四")
	case 'e':
		fmt.Println("周五")
	case 'f':
		fmt.Println("周六")
	case 'g':
		fmt.Println("周日")
	case 'h', 'j', 'k': // 可以有多个
		fmt.Println("其他")
	default:
		fmt.Println("输入有误")
	}
```

### `switch fallthrough`

```go
switch { // swithch 后面可以不写，类似于 if else 分支
case key == 'a':
	fmt.Println("a")
	fallthrough // 穿透，会接着执行后面的一个 case 并且不用判断条件是否成立
case key == 'b':
	fmt.Println("b")
	// fallthrough
	// case key == 'c':
	// 	fmt.Println("c")
}
```

### `for` 循环

```go
for i := 1; i <= 10; i++ {
	fmt.Println("hello world")
}
```

### `for` 死循环

```go
flag := 1
for { // 不加条件
	fmt.Println(flag)
	flag++
	if flag > 100 {
		break
	}
}
```

### `for` 实现 while 循环

```go
flag2 := 100
for flag2 >= 0 {
	fmt.Println(flag2)
	flag2--
}
```

### `goto`

```go
fmt.Println("hello world1")
goto lable1 // 执行到该位置，跳转到 label1 位置接着执行
fmt.Println("hello world2")
fmt.Println("hello world3")
lable1:
fmt.Println("hello world4")
fmt.Println("hello world5")
fmt.Println("hello world6")
```

## 字符串操作

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 1. 字符串长度
	str := "hello 世界"
	fmt.Println(len(str)) // 12 ,ascii 字符占用 1 字节，汉子占用 3 个字节

	// 2. 字符串遍历
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // hello ä¸ç
	}
	fmt.Println()

	// 转换 防止乱码
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i]) // hello 世界
	}
	fmt.Println()

	// 3. 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println(n)
	}

	// 4. 整数转字符串
	nStr := strconv.Itoa(123456)
	fmt.Printf("%T,%s \n", nStr, nStr) //string,123456

	// 5. 字符串转 []byte
	var bytes = []byte("hello")
	fmt.Println(bytes) //[104 101 108 108 111]

	// 6. 把 []byte 转　字符串
	bStr := string([]byte{97, 98, 99})
	fmt.Println(bStr) //abc

	// 7. 十进制转 二、八、十六

	rStr := strconv.FormatInt(123, 2)
	fmt.Println("123对应二进制：", rStr) // 1111011
	rStr = strconv.FormatInt(123, 8)
	fmt.Println("123对应八进制：", rStr) // 173
	rStr = strconv.FormatInt(123, 16)
	fmt.Println("123对应十六进制：", rStr) // 7b

	// 8. 查找子串
	fmt.Println(strings.Contains("hello world", "hello")) // true
	fmt.Println(strings.Contains("hello world", "lol"))   // false

	// 9. 返回不重复子串
	fmt.Println(strings.Count("ceheses", "e")) // 3
	fmt.Println(strings.Count("ceheses", "a")) // 0

	// 10. 字符串比较(==,区分大小写),不区分大小写如下
	fmt.Println(strings.EqualFold("abc", "ABc")) // tue
	fmt.Println("abc" == "ABC")                  // false

	// 11. 返回子串第一次出现的索引
	fmt.Println(strings.Index("abc_hello", "hello"))     // 4
	fmt.Println(strings.Index("abc_hello", "z"))         // -1
	fmt.Println(strings.LastIndex("abc abc abc", "abc")) // 8

	// 12. 子串替换，返回新串，1　表示替换第一个，-1 表示替换所有
	fmt.Println(strings.Replace("abc abc", "abc", "xyz", 1))  // xyz abc
	fmt.Println(strings.Replace("abc abc", "abc", "xyz", -1)) // xyz xyz

	// 13. 字符串拆分字符数组
	fmt.Println(strings.Split("hello,world,go", ",")) // [hello world go]

	// 14. 大小写转换
	fmt.Println(strings.ToUpper("Hello")) // HELLO
	fmt.Println(strings.ToLower("HELLO")) // hello

	// 15. 去掉左右两边空格
	fmt.Printf("%q\n", strings.TrimSpace(" hello world   ")) //"hello world"

	// 16. 指定去掉两边空格
	fmt.Println(strings.Trim(" ! hel!lo! ", " !")) //hel!lo

	// 17. 前缀，后缀
	fmt.Println(strings.HasPrefix("helloworld", "he")) // true
	fmt.Println(strings.HasSuffix("helloworld", "d"))  // true

}
```
