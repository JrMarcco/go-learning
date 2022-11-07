package rpc

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloServiceInf interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInf) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct {
	conn    net.Conn
	isLogin bool
}

// Hello
// 满足 Go 语言的 RPC 规则：
//
//	1、方法只能有两个可序列化的参数且第二个参数是指针类型
//	2、返回一个 error 类型
func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func (h *HelloService) HelloV2(request string, reply *string) error {
	if h.isLogin {
		*reply = "hello: " + request + ", from " + h.conn.RemoteAddr().String()
		return nil
	}
	return errors.New("please login first")
}

func (h *HelloService) Login(request string, _ *string) error {
	if request != "user:password" {
		return errors.New("auth failed")
	}
	log.Println("Login Success")
	h.isLogin = true
	return nil
}
