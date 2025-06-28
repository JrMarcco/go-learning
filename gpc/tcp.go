package gpc

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

const lenBytes = 8

func ReadMsg(conn net.Conn) (bs []byte, err error) {
	msgLenBytes := make([]byte, lenBytes)
	length, err := conn.Read(msgLenBytes)

	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%v", msg)
		}
	}()
	if err != nil {
		return nil, err
	}

	if length != lenBytes {
		return nil, errors.New("read msg length error")
	}

	dataLen := binary.BigEndian.Uint64(msgLenBytes)
	bs = make([]byte, dataLen)

	_, err = io.ReadFull(conn, bs)
	return bs, err
}

func EncodeMsg(bs []byte) []byte {
	encode := make([]byte, lenBytes+len(bs))
	binary.BigEndian.PutUint64(encode[:lenBytes], uint64(len(bs)))
	copy(encode[lenBytes:], bs)
	return encode
}
