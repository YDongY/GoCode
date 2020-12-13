package process

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/07_网络编程/TCPProject/server/utils"
)

type SmsProcess struct {
}

func (s *SmsProcess) SendGroup(content string) (err error) {
	// 1. 创建一个 mes

	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.User.UserId = CurUser.UserId
	smsMes.User.UserStatus = CurUser.UserStatus

	// 2. 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroup json.Marshal err:", err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroup json.Marshal err:", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)

	if err != nil {
		fmt.Println("SendGroup WritePkg err:", err)
		return
	}
	return
}
