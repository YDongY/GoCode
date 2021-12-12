package testing

import "testing"

func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	outputs := [...]int{1, 4, 9}

	for i, v := range inputs {
		ret := Square(v)
		if outputs[i] != ret {
			t.Errorf("input is %d, the expected is %d, the actual %d", v, outputs[i], ret)
		}
	}

	// t.Fail() / t.Error() 测试失败，剩余代码继续执行，其他测试继续执行
	// t.FailNow() / t.Fatal() 测试失败，剩余代码不执行，其他测试继续执行

	/**
		```bash
		$ go test -v -cover
		```

		断言包：https://github.com/stretchr/testify
	 */
}
