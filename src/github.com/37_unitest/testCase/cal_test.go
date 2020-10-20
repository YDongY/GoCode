package main

import "testing"

// 必须 Test 开头，且后缀必须大写字母开头,且当前文件必须test_开头，有助于 go test 查找
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10)执行错误")
	}
	t.Logf("AddUppder(10)执行正确")
}
