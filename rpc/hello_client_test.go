package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestHelloClientV1(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "v1", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}

func TestHelloClientV2(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	err = client.Call(HelloServiceName+".Hello", "v2", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInf = (*HelloServiceClient)(nil)

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName+".Hello", request, reply)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{
		Client: client,
	}, nil
}

func TestHelloClientV3(t *testing.T) {
	client, err := DialHelloService("tcp", ":8081")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	var reply string
	err = client.Hello("v3", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}

func TestHelloClientV4(t *testing.T) {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal("dialing", err)
	}

	// 基于 json 编码
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "v4", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}

func TestHelloClientV5(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	hc := client.Go("HelloService.Hello", "v5", new(string), nil)
	hc = <-hc.Done
	if err = hc.Error; err != nil {
		log.Fatal(err)
	}

	log.Printf("args: %s\n", hc.Args.(string))
	log.Printf("reply: %s\n", *hc.Reply.(*string))
}

func TestHelloClientV6(t *testing.T) {
	client, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	err = client.Call("HelloService.HelloV2", "v6", &reply)
	if err != nil {
		log.Println(err)

		err = client.Call("HelloService.Login", "user:password", &reply)
		if err != nil {
			log.Println(err)
		}

		err = client.Call("HelloService.HelloV2", "v6", &reply)

		log.Printf("after login: %s\n", reply)
		return
	}

	log.Printf("before login: %s\n", reply)
}
