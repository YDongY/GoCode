package processes

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/07_网络编程/TCPProject/server/model"
	"github.com/07_网络编程/TCPProject/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (u *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历 UserMgr
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		// 开始通知
		up.NotifyOnline(userId)
	}
}

func (u *UserProcess) NotifyOnline(userId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("NotifyOnline json.Marshal err:", err)
		return
	}
	// 序列化后的赋值给 Data
	mes.Data = string(data)

	// 再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("NotifyOnline json.Marshal err:", err)
		return
	}

	// 发送
	tf := &utils.Transfer{
		Conn: u.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyOnline WritePkg err:", err)
	}
}

func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal LoginMes err:", err)
		return
	}
	// 1. resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	// 2. 声明一个 LoginResMes
	var registerResMes message.RegisterResMes
	// 3. 数据库完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ErrorUserExists {
			registerResMes.Code = 505
			registerResMes.Error = model.ErrorUserExists.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal loginResMes err:", err)
	}

	resMes.Data = string(data)

	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal resMes err:", err)
		return
	}

	tf := &utils.Transfer{Conn: u.Conn}
	err = tf.WritePkg(data)
	return
}

func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 1. 从 mes 中取出 msg.data
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal LoginMes err:", err)
		return
	}

	// 1. resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	// 2. 声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ErrorUserNotExists {
			// 不合法
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ErrorUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务内部错误"
		}
	} else {
		// 合法
		loginResMes.Code = 200
		// 登录成功添加到 map
		u.UserId = loginMes.UserId
		userMgr.AddOnlineUser(u)
		// 上线通知
		u.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户 id 放入到 UsersId
		for id, _ := range userMgr.onlineUsers {
			if id == loginMes.UserId {
				continue
			}
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Printf("%v 登录成功\n", user)
	}

	// 3. 将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal loginResMes err:", err)
	}

	// 4. 赋值 data 给 resMes
	resMes.Data = string(data)

	// 5. 对 resMes 序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal resMes err:", err)
		return
	}
	// 6. 发送 data

	// 使用了结构体
	tf := &utils.Transfer{Conn: u.Conn}
	err = tf.WritePkg(data)
	return
}
