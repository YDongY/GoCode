package main

import (
	"encoding/json"
	"fmt"
)

type monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	var m = monster{
		Name:     "tom",
		Age:      22,
		Birthday: "2020-11-11",
		Sal:      2000,
		Skill:    "running",
	}

	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println(string(data))
}

func testMap() {
	var map1 map[string]interface{}
	map1 = make(map[string]interface{})
	map1["name"] = "jack"
	map1["age"] = 22
	map1["sal"] = 2000

	data, err := json.Marshal(map1)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println(string(data))
}

func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["nam1"] = "tom"
	m1["age"] = 22
	m1["address"] = [2]string{"北京", "上海"}
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["nam1"] = "jack"
	m2["age"] = 24
	m2["address"] = [2]string{"广州", "深圳"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Println(string(data))
}

func main() {

	testStruct()
	testMap()
	testSlice()
}
