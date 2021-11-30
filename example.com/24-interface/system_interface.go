package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (person *Person) String() string { // 类似于 Python 中的　__str__，方法名称，方法必须是 String
	return fmt.Sprintf("%s-%d", person.Name, person.Age)
}

// Write(p []byte) (n int, err error)
func (person *Person) Write(p []byte) (n int, err error) {
	return
}

// Read(p []byte) (n int, err error)
func (person *Person) Read(p []byte) (n int, err error) {
	return
}

// Close() error
func (person *Person) Close() error {
	return nil
}

func main() {
	p := Person{Name: "jack", Age: 20}
	fmt.Println(&p)

	// Go 语言是面向接口，任何一个结构体类型实现了某些接口，则该接口体就具备某些功能
	// 类似于 Python 中的协议，实现特定的协议，类就可以拥有特定是功能
}