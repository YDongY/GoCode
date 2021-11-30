package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalf("Dial error:%v\n", err)
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Fatalf("Write error:%v\n", err)
		}
	}
}
