package main

import "fmt"

func main() {
	// if-else

	age := 10

	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	// if-else if-else
	var num int
	fmt.Println("请输入num值:")
	fmt.Scanln(&num)

	if num >= 100 && num < 1000 {
		fmt.Println("num>=100")
	} else if num >= 1000 {
		fmt.Println("num>=1000")
	} else {
		fmt.Println("num<100")
	}

	// switch 不同添加 break 自动在结尾添加 break
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

	// for

	for i := 1; i <= 10; i++ {
		fmt.Println("hello world")
	}

	// 死循环
	flag := 1
	for {
		fmt.Println(flag)
		flag++
		if flag > 100 {
			break
		}
	}

	// 相当于 while 循环

	flag2 := 100

	for flag2 >= 0 {
		fmt.Println(flag2)
		flag2--
	}

	// 遍历字符串
	str := "helloworld"
	for i := 0; i < len(str); i++ { // 按照字节方式遍历
		fmt.Println(string(str[i]))
	}

	for _, val := range str { // 按照字符遍历
		fmt.Println(string(val))
	}

	// goto

	fmt.Println("hello world1")
	goto lable1
	fmt.Println("hello world2")
	fmt.Println("hello world3")
lable1:
	fmt.Println("hello world4")
	fmt.Println("hello world5")
	fmt.Println("hello world6")
}
