package rpc

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestHelloServerV1(t *testing.T) {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatal("Register error: ", err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}

	rpc.ServeConn(conn)
}

func TestHelloServerV2(t *testing.T) {
	if err := RegisterHelloService(new(HelloService)); err != nil {
		log.Fatal("Register error: ", err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}

func TestHelloServerV3(t *testing.T) {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatal("Register error: ", err)
	}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		// 基于 json 编码
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

// 在 http 协议上提供 json rpc 服务
func TestHelloServerV4(t *testing.T) {
	if err := rpc.RegisterName("HelloService", new(HelloService)); err != nil {
		log.Fatal("Register error: ", err)
	}

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}

		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

// Context
func TestHelloServerV5(t *testing.T) {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go func() {
			defer func() {
				_ = conn.Close()
			}()

			p := rpc.NewServer()
			_ = p.Register(&HelloService{conn: conn})
			p.ServeConn(conn)
		}()
	}
}
