package model

import (
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"github.com/gomodule/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (u *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ErrorUserNotExists
		}
		return
	}
	user = &message.User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("getUserById json.Unmarshal err", err)
		return
	}
	return
}

func (u *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	conn := u.pool.Get() // Redis 连接池
	defer conn.Close()
	user, err = u.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ErrorUserPwd
	}
	return
}

func (u *UserDao) Register(user *message.User) (err error) {
	conn := u.pool.Get() // Redis 连接池
	defer conn.Close()
	_, err = u.getUserById(conn, user.UserId)
	if err == nil {
		err = ErrorUserExists
		return
	}

	// 说明用户没有注册过
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误　err=", err)
		return
	}
	return
}
