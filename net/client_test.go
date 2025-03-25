package net

import (
	"encoding/binary"
	"log"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	conn, err := net.DialTimeout("tcp", ":8081", time.Second)
	if err != nil {
		t.Fatal(err)
	}

	msg := "hello from client"
	msgLen := len(msg)

	msgLenBs := make([]byte, 8)
	binary.LittleEndian.PutUint64(msgLenBs, uint64(msgLen))
	data := append(msgLenBs, []byte(msg)...)

	_, err = conn.Write(data)
	if err != nil {
		if err = conn.Close(); err != nil {
			log.Println(err)
		}
		return
	}

	respBs := make([]byte, 16)
	_, err = conn.Read(respBs)
	if err != nil {
		if err = conn.Close(); err != nil {
			log.Println(err)
		}
		return
	}

	log.Println(string(respBs))
}
