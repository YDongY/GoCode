package main

import (
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("open log file failed, err:", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[logger]")
	log.Println("这是一条普通的日志")

	// log.Printf("这是一条%s的日志\n", "普通")
	// log.Fatalln("这是一条会触发 fatal 的日志")
	// log.Panicln("这是一条会触发 panic 的日志")

	// logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。
	// Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。

	logger := log.New(os.Stdout, "[logger]", log.Llongfile|log.Lmicroseconds|log.Ldate)
	logger.Println("这是一条普通的日志")
}
