package _3_escape

import (
	"testing"
)

func TestEscape(t *testing.T) {
	t.Log("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	t.Logf("%d %x %X %o %d \n", 18, 11, 12, 18, 18) // 18 b C 22 18
	t.Logf("%f \n", 1.5)                            // 1.500000
	t.Logf("%t \n", false)
	t.Logf("%c \n", 97)            // a
	t.Logf("%s \n", "hello world") // hello world
	t.Logf("%q \n", "hello world") // "hello world"
	t.Logf("%v \n", "你好")          // 你好
	t.Logf("%T \n", false)         // bool

	i := 5
	t.Logf("%p \n", &i)   // 0xc00001a088
	t.Logf("%d%% \n", 20) // 20%
}
