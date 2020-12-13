package main

import (
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/07_网络编程/TCPProject/server/processes"
	"github.com/07_网络编程/TCPProject/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		sp := &processes.SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (p *Processor) process2() (err error) {
	// 读取客户端发送的信息
	for {
		// 读取数据
		tf := &utils.Transfer{Conn: p.Conn}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出")
				return err
			} else {
				fmt.Println("readPkg err:", err)
				return err
			}
		}
		fmt.Println("mes=", mes)
		err = p.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
