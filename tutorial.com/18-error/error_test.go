package _8_error

import (
	"errors"
	"fmt"
	"testing"
)

func readFile(name string) error {
	if name == "config.ini" {
		return nil
	} else {
		// 自定义错误 errors.New("错误说明")
		return errors.New("读取文件错误")
	}
}

func TestError(t *testing.T) {
	/**
	  1. 没有异常机制
	  2. error 类型实现了 error 接口
	  3. errors.New 来创建一个错误实例
	*/

	defer func() {
		if err := recover(); err != nil {
			// recover 错误恢复，err 就是 panic 传递过来的
			fmt.Println("err = ", err) // err =  读取文件错误
		}
	}()

	if err := readFile("config.ini2"); err != nil {
		panic(err) // 程序终止，执行 defer
	}

	fmt.Println("Done")
}
