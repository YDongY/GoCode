package main

import (
	"testing"
)

func TestIfElse(t *testing.T) {
	// ------------------ if-else ------------------

	age := 10

	if age >= 18 {
		t.Log("已成年")
	} else {
		t.Log("未成年")
	}

	if a := 10 * 2; a > age {
		t.Log("a>10")
	} else {
		t.Log("a<10")
	}
}

func TestIfElseIf(t *testing.T) {
	// ------------------ if - else if - else ------------------
	var num int = 998

	if num >= 100 && num < 1000 {
		t.Log("num>=100")
	} else if num >= 1000 {
		t.Log("num>=1000")
	} else {
		t.Log("num<100")
	}
}

func TestSwitch(t *testing.T) {
	// ------------------ switch ------------------
	// switch 不用添加 break 自动在结尾添加 break
	var key byte = 'c'
	switch key { // key 常量、变量、返回值
	// case 类型必须与 switch 一致，case 表达式为常量，不能重复
	case 'a':
		t.Log("周一")
	case 'b':
		t.Log("周二")
	case 'c':
		t.Log("周三")
	case 'd':
		t.Log("周四")
	case 'e':
		t.Log("周五")
	case 'f':
		t.Log("周六")
	case 'g':
		t.Log("周日")
	case 'h', 'j', 'k': // 可以有多个
		t.Log("其他")
	default:
		t.Log("输入有误")
	}
}

func TestSwitchFallthrough(t *testing.T) {
	var key byte = 'b'
	switch { // switch 后面可以不写，类似于 if else 分支
	case key == 'a':
		t.Log("a")
		fallthrough // 穿透，会接着执行后面的一个 case 并且不用判断条件是否成立
	case key == 'b':
		t.Log("b")
		fallthrough
	case key == 'c':
		t.Log("c")
	}
}

func TestFor(t *testing.T) {
	// ------------------ for i 循环 ------------------

	for i := 1; i <= 10; i++ {
		t.Log("hello world")
	}

	// ------------------ for 死循环 ------------------
	flag := 1
	for {
		t.Log(flag)
		flag++
		if flag > 10 {
			break
		}
	}

	// ------------------ for while 循环 ------------------

	flag2 := 10

	for flag2 >= 0 {
		t.Log(flag2)
		flag2--
	}

	// ------------------ for-i vs for-r ------------------
	str := "hello world"
	for i := 0; i < len(str); i++ { // 按照字节方式遍历
		t.Log(string(str[i]))
	}

	for _, val := range str { // 按照字符遍历
		t.Log(string(val))
	}
}

func TestGoto(t *testing.T) {
	// ------------------ goto ------------------

	t.Log("hello world1")
	goto label
	t.Log("hello world2")
	t.Log("hello world3")
label:
	t.Log("hello world4")
	t.Log("hello world5")
	t.Log("hello world6")
}
