package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/07_网络编程/TCPProject/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (t *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)

	// 发送长度
	n, err := t.Conn.Write(t.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write len err:", err)
		return
	}

	// 发送 data 本身
	n, err = t.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write pkg err:", err)
		return
	}
	return
}

func (t *Transfer) ReadPkg() (mes message.Message, err error) {
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read len err:", err)
		return
	}

	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[:4])
	n, err := t.Conn.Read(t.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn.Read pkg err:", err)
		return
	}

	// pkgLen 反序列化 -> message.Message
	err = json.Unmarshal(t.Buf[:pkgLen], &mes)
	if err != nil {
		// fmt.Println("json.Unmarshal err:", err)
		return
	}
	return
}
