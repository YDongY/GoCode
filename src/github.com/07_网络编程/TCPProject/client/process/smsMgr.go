package process

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
)

func outputGroupMes(mes *message.Message) {
	// 1. 反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal err:", err)
		return
	}
	// 显示信息
	fmt.Printf("用户id:\t%d 对大家说:\t%s", smsMes.User.UserId, smsMes.Content)
	fmt.Println()
}
