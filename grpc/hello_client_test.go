package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestHelloClientV1(t *testing.T) {
	conn, err := grpc.Dial(
		":8081", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "v1"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reply.GetValue())
}

func TestHelloClientV2(t *testing.T) {
	conn, err := grpc.Dial(
		":8081", grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := NewHelloServiceClient(conn)

	// 获取返回的流对象
	stream, err := client.Channel(context.Background())

	// 单独 goroutine 执行发送操作
	go func() {
		for {
			if err := stream.Send(&String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		if reply, err := stream.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		} else {
			log.Printf("receive msg from server: %s\n", reply.GetValue())
		}
	}
}

// 证书认证
func TestHelloClientV3(t *testing.T) {
	creds, err := credentials.NewClientTLSFromFile("./cert/server.pem", "server.grpc.org")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "v3"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reply.GetValue())
}

// 双向证书认证
func TestHelloClientV4(t *testing.T) {
	certificate, err := tls.LoadX509KeyPair("./cert/client.pem", "./cert/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("./cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal(err)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "server.grpc.org", // Subject Alternative Name
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(
		":8081",
		grpc.WithTransportCredentials(creds),
		// 传入 Token 信息
		grpc.WithPerRPCCredentials(&Authentication{
			Username: "jrmarcco",
			Password: "util",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "v4"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply.GetValue())
}
