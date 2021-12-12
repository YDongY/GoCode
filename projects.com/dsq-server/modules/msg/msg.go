package msg

import (
    "github.com/name5566/leaf/network/protobuf"
)


var Processor = protobuf.NewProcessor() // protobuf

func init() {
    Processor.Register(&Number{})
    // Processor.Register(&Number{}) // Json 协议
    // var Processor = json.NewProcessor() // json
}

// Number 一个结构体定义了一个 JSON 消息的格式，消息名为 Number
//type Number struct {
//    Num int
//}
