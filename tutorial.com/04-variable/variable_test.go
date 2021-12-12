package _4_variable

import (
	"testing"
)

func TestVariableInitialValue(t *testing.T) {
	// 变量初始值
	var (
		a int         = 10
		b string      = "mark"
		c []float64   = []float64{3.14}
		d func() bool = func() bool {
			return true
		}
		e struct {
			name string
			age  int
		} = struct {
			name string
			age  int
		}{name: "mark", age: 24}
	)
	t.Log(a, b, c, d(), e) // 0 "" [] <nil> { 0}
}

func TestVariableTypeDeduction(t *testing.T) {
	// 变量类型推断
	var num, name, gender = 10, "jack", true
	t.Log(num, name, gender) // 10 jack true
}

func TestVariableShorter(t *testing.T) {
	// 短变量
	num, name, gender := 10, "jack", true
	t.Log(num, name, gender) // 10 jack true
}
