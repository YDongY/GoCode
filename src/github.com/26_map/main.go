package main

import (
	"fmt"
	"sort"
)

func main() {
	// 声明

	var map1 map[string]string

	fmt.Println(map1) // map[]

	// map1["name"] = "ydony" // panic: assignment to entry in nil map

	// fmt.Println(map1)

	// 使用 map 之前需要 make 分配数据空间
	map1 = make(map[string]string, 10) // 最多存放10
	map1["name"] = "ydongy"
	map1["gender"] = "male"
	fmt.Println(map1) // map[gender:male name:ydongy]

	// 使用方式

	// 方式一
	var map2 map[string]string
	map2 = make(map[string]string, 10)
	fmt.Println(map2)

	// 方式二
	var map3 = make(map[string]string, 10)
	fmt.Println(map3)

	// 方式三
	map4 := make(map[string]string)

	map4["name"] = "tom"
	map4["city"] = "北京"
	map4["gender"] = "male"
	fmt.Println(map4) // map[city:北京 gender:male name:tom]

	// 方式四

	map5 := map[string]string{
		"name": "jack",
		"age":  "20",
	}

	fmt.Println(map5) // map[age:20 name:jack]

	// map crud
	map6 := make(map[string]string, 5)

	map6["name"] = "jack"
	map6["age"] = "23"
	fmt.Println(map6)

	delete(map6, "name") // 删除 key

	fmt.Println(map6)

	// 清空map　重新创建空间
	// map6 = make(map[string]string)

	fmt.Println(map6)

	// 查找

	val, ok := map6["age"]
	if ok {
		fmt.Println("值=", val)
	} else {
		fmt.Println("没有值:", ok)
	}

	// map　遍历

	map7 := map[string]string{
		"name":   "jack",
		"age":    "23",
		"gender": "male",
		"city":   "上海",
	}

	for k, v := range map7 {
		fmt.Println(k, v)
	}

	// map len

	fmt.Println(len(map7)) // 4

	// map 切片

	var mapSlice []map[string]string

	mapSlice = make([]map[string]string, 2) // 切片大小，2个map

	mapSlice[0] = make(map[string]string)
	mapSlice[1] = make(map[string]string)

	mapSlice[0]["name"] = "jack"
	mapSlice[0]["age"] = "23"
	mapSlice[1]["name"] = "tom"
	mapSlice[1]["age"] = "26"

	fmt.Println(mapSlice) // [map[age:23 name:jack] map[age:26 name:tom]]

	// append 添加 map 切片

	newMap := map[string]string{
		"name": "ydony",
		"age":  "23",
	}

	mapSlice = append(mapSlice, newMap)

	fmt.Println(mapSlice) // [map[age:23 name:jack] map[age:26 name:tom] map[age:23 name:ydony]]

	// map 排序

	// 1. 先将map的key 存放到切片
	// 2. 对切片排序
	// 3. 遍历切片

	map8 := map[int]int{
		1: 100,
		2: 200,
		4: 400,
		3: 300,
	}

	fmt.Println(map8) // map[1:100 2:200 3:300 4:400]

	var keys []int

	for k, _ := range map8 {
		keys = append(keys, k)
	}

	sort.Ints(keys) // [1 2 3 4]

	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(map8[k])
	}

}
