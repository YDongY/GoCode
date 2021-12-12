package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// ----------- 声明 -----------
	var m map[string]string
	t.Log(m) // map[]
	// m["name"] = "mark" // panic: assignment to entry in nil map

	// ----------- 使用 map 之前需要 make 分配数据空间 -----------
	m = make(map[string]string, 10) // 最多存放10
	m["name"] = "mark"
	m["gender"] = "male"
	t.Log(m) // map[gender:male name:mark]
}

func TestMapInit(t *testing.T) {
	// 方式一
	var m1 map[string]string
	m1 = make(map[string]string, 10)
	t.Log(m1, len(m1))

	// 方式二
	var m2 = make(map[string]string, 10)
	t.Log(m2, len(m2))

	// 方式三
	m3 := make(map[string]string)
	m3["name"] = "tom"
	m3["city"] = "北京"
	m3["gender"] = "male"
	t.Log(m3, len(m3)) // map[city:北京 gender:male name:tom]

	// 方式四
	m4 := map[string]string{
		"name": "jack",
		"age":  "20",
	}
	t.Log(m4, len(m4)) // map[age:20 name:jack]
}

func TestMapCRUD(t *testing.T) {
	// 创建
	m := map[string]string{
		"name": "jack",
		"age":  "20",
	}

	// 删除 key
	delete(m, "name")
	fmt.Println(m)

	// 根据 key 查找 value
	// 当访问的 Key 不存在的时候，仍会返回零值，所以如果某个 value 刚好等于零值，则不能通过 nil 来判断元素是否存在
	if val, ok := m["age"]; ok {
		fmt.Println("值=", val)
	} else {
		fmt.Println("没有值:", ok)
	}

	// 清空 map　重新创建空间
	m = make(map[string]string)
}

func TestMapTravel(t *testing.T) {
	// map　遍历
	m := map[string]string{
		"name":   "jack",
		"age":    "23",
		"gender": "male",
		"city":   "上海",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func TestMapValueType(t *testing.T) {
	funcMap := map[string]func(op int) int{}
	{
	}
	funcMap["func1"] = func(op int) int {
		return op
	}

	funcMap["func2"] = func(op int) int {
		return op * op
	}

	funcMap["func3"] = func(op int) int {
		return op * op * op
	}

	t.Log(funcMap["func1"](2), funcMap["func2"](2), funcMap["func3"](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing\n", n)
	} else {
		t.Logf("%d is not existing\n", n)
	}
}
