package _2_protobuf

import (
	"github.com/golang/protobuf/proto"
	"golang-dev/tutorial.com/42-protobuf/hello"
	"testing"
)

func TestHelloProto(t *testing.T) {
	h := &hello.Hello{
		Name: "Mark",
		Age:  24,
	}

	// 序列化
	marshal, err := proto.Marshal(h)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("binary data:", marshal)

	// 反序列化
	msgEntity := &hello.Hello{}
	if err = proto.Unmarshal(marshal, msgEntity); err != nil {
		t.Fatal(err)
	}
	t.Log(msgEntity.Name, msgEntity.Age)           // Mark 24
	t.Log(msgEntity.GetName(), msgEntity.GetAge()) // Mark 24
	t.Log(msgEntity.String())                      // Name:"Mark" Age:24
}
