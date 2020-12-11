package main

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

type Config struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func testReadFile() {
	// 反序列化读取配置
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	c := Config{}
	err = json.Unmarshal(f, &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v \n", c)
}

func testTOML() {
	// 使用 TOML 文件
	// TOML ( Tom’s Obvious, Minimal Language ）是一种专为存储配置文件而设计的格式
	// go get github.com/BurntSushi/toml
	c := Config{}
	_, err := toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v \n", c)
}

func main() {
	testTOML()
}
