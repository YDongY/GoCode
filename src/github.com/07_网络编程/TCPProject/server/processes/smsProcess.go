package processes

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/07_网络编程/TCPProject/server/utils"
	"net"
)

type SmsProcess struct {
}

func (s *SmsProcess) SendGroupMes(mes *message.Message) {
	// 遍历 map，转发消息

	// 1. 取出内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal err:", err)
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal err:", err)
	}

	for id, up := range userMgr.onlineUsers {
		if smsMes.User.UserId == id {
			continue
		}
		s.SendMesToEach(data, up.Conn)
	}
}

func (s *SmsProcess) SendMesToEach(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEach WritePkg err:", err)
	}
}
