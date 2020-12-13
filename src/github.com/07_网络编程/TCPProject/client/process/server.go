package process

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/07_网络编程/TCPProject/server/utils"
	"net"
	"os"
)

func ShowMenu() {
	fmt.Println("----------登录成功----------")
	fmt.Println("\t 1. 显示用户在线列表")
	fmt.Println("\t 2. 发送消息")
	fmt.Println("\t 3. 信息列表")
	fmt.Println("\t 4. 退出系统")
	fmt.Println("\t 请选择【1-4】")

	var key int
	var content string

	smsProcess := &SmsProcess{}

	fmt.Scanf("%d \n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入>>>:")
		fmt.Scanln(&content)
		smsProcess.SendGroup(content)
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("已经退出系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的有误")
	}
}

func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{Conn: conn}

	for {
		// 客户端读取
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() error", err)
			return
		}
		// fmt.Printf("mes=%v \n", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType: // 上线通知
			// 1. 取出消息
			var notifyUserStatusMes message.NotifyUserStatusMes
			err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			if err != nil {
				fmt.Println("serverProcessMes json.Unmarshal err:", err)
			}
			// 2. 把上线用户加入到 map
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsMesType: // 群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("消息类型无法识别")
		}
	}
}
