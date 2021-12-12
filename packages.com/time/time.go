package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// time.Time,2021-11-30 14:09:10.7117966 +0800 CST m=+0.002213001
	fmt.Printf("%T,%v \n", now, now)

	// 获取日期
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, int(month), day, hour, minute, second) // 2021 November 11 30 14 11 43

	// 格式化
	fmt.Println(now.Format("2006/01/02 15:04:05")) // 2021/11/30 14:12:00
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2021-11-30 14:12:00
	fmt.Println(now.Format("2006-01-02"))          // 2021-11-30
	fmt.Println(now.Format("01"))                  // 11
	fmt.Println(now.Format("2006"))                // 2021

	// 解析字符串格式时间
	// 加载时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	localTime, _ := time.ParseInLocation("2006/01/02 15:04:05", "2021/11/30 14:12:00", local)
	fmt.Println(localTime)

	// 时间戳
	timestamp1 := now.Unix()     // 时间戳 1638252940
	timestamp2 := now.UnixNano() // 纳秒时间戳 1638252940792590900
	fmt.Println(timestamp1, timestamp2)

	// 时间戳转换为时间格式
	t := time.Unix(now.Unix(), 0)

	// time.Time,2021-11-30 14:17:30 +0800 CST
	fmt.Printf("%T,%v \n", t, t)

	// 时间常量定义
	//const (
	//	Nanosecond  Duration = 1
	//	Microsecond          = 1000 * Nanosecond
	//	Millisecond          = 1000 * Microsecond
	//	Second               = 1000 * Millisecond
	//	Minute               = 60 * Second
	//	Hour                 = 60 * Minute
	//)

	// 时间操作
	hourAfter := now.Add(time.Hour) // 当前时间加1小时后的时间
	// time.Time,2021-11-30 15:20:45.41066 +0800 CST m=+3600.002217001
	fmt.Printf("%T,%v \n", hourAfter, hourAfter)

	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("两个时间之差:", end.Sub(start))

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(secondsEastOfUTC, beijing)

	// 同一时刻，但是不同时区
	d1 := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	d2 := time.Date(2000, 2, 1, 20, 30, 0, 0, beijing)

	fmt.Println(d1 == d2)     // False
	fmt.Println(d1.Equal(d2)) // True，Equal 可以比较同一时刻，但是不同时区的时间

	// ------ Before ------
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	fmt.Printf("year2000.Before(year3000) = %v\n", year2000.Before(year3000)) // True
	fmt.Printf("year3000.Before(year2000) = %v\n", year3000.Before(year2000)) // False

	// ------ After ------
	year2000 = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 = time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	fmt.Printf("year3000.After(year2000) = %v\n", year3000.After(year2000)) // True
	fmt.Printf("year2000.After(year3000) = %v\n", year2000.After(year3000)) // False

	// 定时器
	ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
	for t := range ticker {
		fmt.Println(t) // 每 1s 执行一次
	}
}
