package internal

import (
    "dsq-server/modules/msg"
    "github.com/name5566/leaf/gate"
    "github.com/name5566/leaf/log"
    "google.golang.org/protobuf/proto"
    "reflect"
)

func init() {
    // 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleGuessNumber
    handler(&msg.Number{}, handleGuessNumber)
}

func handler(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleGuessNumber(args []interface{}) {
    // 收到的 Number 消息
    m := args[0].(*msg.Number)
    // 消息的发送者
    a := args[1].(gate.Agent)

    // ---------------------- Json ------------------------
    // 输出收到的消息的内容
    //log.Debug("number: %v", m.Num)
    //
    //// 给发送者回应一个 Hello 消息
    //a.WriteMsg(&msg.Number{
    //    Num: 1,
    //})

    // --------------------- ProtoBuf --------------------------
    // 输出收到的消息的内容
    log.Debug("number %v", m.GetNum())

    // 给发送者回应一个 Hello 消息
    a.WriteMsg(&msg.Number{
        Num: *proto.Uint32(32),
    })
}
