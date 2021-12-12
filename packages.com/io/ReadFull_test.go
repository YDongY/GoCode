package io

import (
	"encoding/binary"
	"io"
	"strings"
	"testing"
)

func TestReadFull(t *testing.T) {
	var b [4]byte
	bufMsgLen := b[:2]

	reader := strings.NewReader("hello")
	if n, err := io.ReadFull(reader, bufMsgLen); err != nil {
		t.Log(err)
	} else {
		t.Log(n)
	}

	msgLen := uint32(binary.BigEndian.Uint16(bufMsgLen))

	msgData := make([]byte, msgLen)
	if n, err := io.ReadFull(reader, msgData); err != nil {
		t.Log(err)
	} else {
		t.Log(n)
	}
}
