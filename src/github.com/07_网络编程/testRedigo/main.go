package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 最大连接数
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { //　初始化连接代码
			return redis.Dial("tcp", ":6379", redis.DialDatabase(1))
		},
	}
}

func testString() {
	// 连接 redis
	c, err := redis.Dial("tcp", ":6379", redis.DialDatabase(1))
	if err != nil {
		// handle error
		fmt.Println("error redis.Dial")
		return
	}
	fmt.Println("conn redis success...")
	defer c.Close()

	// 向 redis 写入数据
	_, err = c.Do("SET", "name", "jack")
	if err != nil {
		fmt.Println("do set error")
	}

	// 读取数据
	s, err := redis.String(c.Do("GET", "hello"))
	if err != nil {
		fmt.Println("do get error")
	}
	fmt.Printf("%#v\n", s)

	// 批量设置
	_, err = c.Do("MSET", "age", 20, "score", 99.9)
	if err != nil {
		fmt.Println("do mset error")
	}

	r, err := redis.Strings(c.Do("MGET", "age", "name"))
	if err != nil {
		fmt.Println("do mget error")
	}
	for i, v := range r {
		fmt.Printf("%d %s \n", i, v)
	}
}

func testPool() {
	// Redis 连接池
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom猫猫")
	if err != nil {
		fmt.Println("do set error")
		return
	}

	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("do get error")
	}
	fmt.Printf("%#v \n", r)
}

func main() {
	testPool()
}
