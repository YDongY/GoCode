package gate

import (
    "dsq-server/modules/game"
    "dsq-server/modules/msg"
)

func init() {
    // 这里指定消息 Number 路由到 game 模块，模块间使用 ChanRPC 通讯，消息路由也不例外
    msg.Processor.SetRouter(&msg.Number{}, game.ChanRPC)
}
