package process

import (
	"fmt"
	"github.com/07_网络编程/TCPProject/client/model"
	"github.com/07_网络编程/TCPProject/common/message"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

func outputOnlineUser() {
	fmt.Println("******当前在线用户列表******")
	for id, user := range onlineUsers {
		fmt.Printf("Id:%d\t Status:%d \n", id, user.UserStatus)
	}
}

func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
