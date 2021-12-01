package main

import (
	"fmt"
	"time"
)

func TimerOnce() {
	// 验证 timer 只能响应1次
	timer2 := time.NewTimer(time.Second)
	for {
		<-timer2.C
		fmt.Println("时间到")
	}
}

func TimerDelay() {
	// timer 实现延时的功能
	time.Sleep(time.Second)
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	<-time.After(2 * time.Second)
	fmt.Println("2秒到")
}

func TimerStop() {
	// 停止定时器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}
}

func TimerRest() {
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

	time.Sleep(10 * time.Second)
}

func Timer() {
	// timer 基本使用
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)
}

func Ticker() {
	// 1.获取 ticker 对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				// 停止
				ticker.Stop()
			}
		}
	}()
	for {
	}
}

func main() {
	Ticker()
}
