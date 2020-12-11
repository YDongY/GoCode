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
	str := "{\"Name\":\"tom\",\"Age\":22,\"Birthday\":\"2020-11-11\",\"Sal\":2000,\"Skill\":\"running\"}"

	var m monster
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)      // {tom 22 2020-11-11 2000 running}
	fmt.Println(m.Name) // tom
}

func testMap() {
	str := "{\"age\":22,\"name\":\"jack\",\"sal\":2000}"
	var map1 map[string]interface{}
	// 反序列化 ,map 不需要 make ，Unmarshal 已经封装了 make
	err := json.Unmarshal([]byte(str), &map1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(map1) // map[age:22 name:jack sal:2000]
}

func testSlice() {
	str := "[{\"address\":[\"北京\",\"上海\"],\"age\":22,\"nam1\":\"tom\"}," +
		"{\"address\":[\"广州\",\"深圳\"],\"age\":24,\"nam1\":\"jack\"}]"

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slice) // [map[address:[北京 上海] age:22 nam1:tom] map[address:[广州 深圳] age:24 nam1:jack]]
}

func main() {
	testStruct()
	testMap()
	testSlice()
}
