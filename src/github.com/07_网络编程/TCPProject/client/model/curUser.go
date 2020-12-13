package model

import (
	"github.com/07_网络编程/TCPProject/common/message"
	"net"
)



type CurUser struct {
	Conn net.Conn
	message.User
}
