package main

import (
    "dsq-server/modules/msg"
    "encoding/binary"
    "fmt"
    "google.golang.org/protobuf/proto"
    "net"
    "testing"
)

func TestJsonMsgServer(t *testing.T) {
    conn, err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
        panic(err)
    }

    // Number 消息（JSON 格式） 对应游戏服务器 Number 消息结构体
    data := []byte(`{
		"Number": {
			"Num": 53
		}
	}`)

    // len + data
    m := make([]byte, 2+len(data))

    // 默认使用大端序
    binary.BigEndian.PutUint16(m, uint16(len(data)))

    copy(m[2:], data)

    // 发送消息
    conn.Write(m)
}

func TestProtoMsgServer(t *testing.T) {
    conn, err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
        panic(err)
    }

    // 这里需要是 msg.pb.go  中的 Number，不能使用自己的 Number struct
    n := &msg.Number{Num: 32}

    fmt.Printf("%T %v\n", n, n)
    // 序列化
    marshal, err := proto.Marshal(n)
    if err != nil {
        t.Fatal(err)
    }

    fmt.Println("marshal:", marshal)

    // 反序列化
    //unmarshal := &msg.Number{}
    //if err = proto.Unmarshal(marshal, unmarshal); err != nil {
    //    t.Fatal(err)
    //}
    //fmt.Println("Unmarshal:", unmarshal)

    // len + id +data
    m := make([]byte, 2+2+len(marshal))

    // 默认使用大端序
    binary.BigEndian.PutUint16(m, uint16(len(marshal)))

    copy(m[4:], marshal)

    // 发送消息
    conn.Write(m)
}
