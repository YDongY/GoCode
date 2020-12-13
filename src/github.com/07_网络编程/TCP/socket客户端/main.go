package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conn to server success...")

	for {
		fmt.Printf("请输入内容：")
		render := bufio.NewReader(os.Stdin)
		content, err := render.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		content = strings.Trim(content, "\n")
		if content == "exit" {
			conn.Close()
			break
		}
		_, err = conn.Write([]byte(content))
		if err != nil {
			log.Fatal(err)
		}
	}

}
