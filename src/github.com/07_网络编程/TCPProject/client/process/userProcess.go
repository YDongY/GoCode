package process

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/07_网络编程/TCPProject/server/utils"
	"net"
)

type UserProcess struct {
}

func (u *UserProcess) Login(userId int, userPwd string) (err error) {
	// fmt.Printf("userID = %d userPwd = %s", userId, userPwd)
	// return

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	defer conn.Close()

	// 发送消息给服务器

	// 1. 创建结构体
	var mes message.Message
	mes.Type = message.LoginMesType

	// 2. 创建一个 LoginMes
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 3. 序列化 LoginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 4. 设置 Data
	mes.Data = string(data)

	// 5. 将 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 这里还需要处理服务器返回的消息
	tf := &utils.Transfer{Conn: conn}

	// 6. 发送 data
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err:", err)
		return
	}

	// 读取数据
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}
	// 将 mes 中的 data 反序列化
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// 初始化 CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		// 显示当前在线列表
		fmt.Println("****** 当前在线列表如下: ******")
		for _, v := range loginResMes.UsersId {
			fmt.Println("用户 id: ", v)

			// 完成用户各自 map 的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()
		// 启动协程，保持和服务通信
		go serverProcessMes(conn)

		// 显示登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (u *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	defer conn.Close()

	// 发送消息给服务器

	// 1. 创建结构体
	var mes message.Message
	mes.Type = message.RegisterMesType
	var registerMes message.RegisterMes

	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	// 赋值数据
	mes.Data = string(data)

	// 将 mes 序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	tf := &utils.Transfer{Conn: conn}

	// 发送 data
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err:", err)
		return
	}

	// 读取数据
	mes, err = tf.ReadPkg() // mes 就是 RegisterResMes
	if err != nil {
		fmt.Println("readPkg err:", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return
}
