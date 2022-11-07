package net

import (
	"encoding/binary"
	"log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func() {
			handle(conn)
		}()
	}
}

func handle(conn net.Conn) {
	for {
		lenBs := make([]byte, 8)
		_, err := conn.Read(lenBs)
		if err != nil {
			if err = conn.Close(); err != nil {
				log.Println(err)
			}
			return
		}

		msgLen := binary.LittleEndian.Uint64(lenBs)
		reqBs := make([]byte, msgLen)

		_, err = conn.Read(reqBs)
		if err != nil {
			if err := conn.Close(); err != nil {
				log.Println(err)
			}
			return
		}

		log.Println(string(reqBs))

		_, err = conn.Write([]byte("hello from server"))
		if err != nil {
			if err = conn.Close(); err != nil {
				log.Println(err)
			}
			return
		}
	}
}
