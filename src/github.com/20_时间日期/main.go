package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T,%v \n", now, now) // time.Time,2020-10-17 22:44:54.647712027 +0800 CST m=+0.000046700

	// 获取日期
	fmt.Println("年：", now.Year())
	fmt.Println("月：", now.Month())
	fmt.Println("月：", int(now.Month()))
	fmt.Println("日：", now.Day())
	fmt.Println("时：", now.Hour())
	fmt.Println("分：", now.Minute())
	fmt.Println("秒：", now.Second())

	// 格式化

	fmt.Println(now.Format("2006/01/02 15:04:05")) // 2020/10/17 22:49:20
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 2020-10-17 22:49:37
	fmt.Println(now.Format("2006-01-02"))          // 2020-10-17
	fmt.Println(now.Format("01"))                  // 10
	fmt.Println(now.Format("2006"))                // 2020

	// 时间常量
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )

	// i := 0
	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond * 100)
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	// 随机数字
	// unix时间戳=1602948366, unixnaon时间戳=1602948366056742753
	fmt.Printf("unix时间戳=%v, unixnaon时间戳=%v", now.Unix(), now.UnixNano())
	
















}
