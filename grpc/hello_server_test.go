package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestHelloServerV1(t *testing.T) {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	_ = grpcServer.Serve(listener)
}

// 证书认证
// 注意需要是用 SANs 扩展
func TestHelloServerV2(t *testing.T) {
	creds, err := credentials.NewServerTLSFromFile("./cert/server.pem", "./cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	_ = grpcServer.Serve(listener)
}

// 双向证书认证
func TestHelloServerV3(t *testing.T) {
	certificate, err := tls.LoadX509KeyPair("./cert/server.pem", "./cert/server.key")
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
		Certificates: []tls.Certificate{certificate}, // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, // 需要并且验证客户端证书
		ClientCAs:    certPool,                       // 客户端证书池
	})

	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		// 拦截器
		grpc.UnaryInterceptor(filter),
	)
	RegisterHelloServiceServer(grpcServer, NewHelloServiceImpl())

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	_ = grpcServer.Serve(listener)
}

func filter(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Println("filter: ", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v\n", r)
		}
	}()
	return handler(ctx, req)
}

func TestWebServerWithGrpc(t *testing.T) {
	mux := http.NewServeMux()

	certificate, err := tls.LoadX509KeyPair("./cert/server.pem", "./cert/server.key")
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
		Certificates: []tls.Certificate{certificate}, // 服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, // 需要并且验证客户端证书
		ClientCAs:    certPool,                       // 客户端证书池
	})

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	_ = http.ListenAndServeTLS(":8080", "./cert/server.pem", "./cert/server.key",
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.ProtoMajor != 2 {
				mux.ServeHTTP(writer, request)
				return
			}

			if strings.Contains(request.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(writer, request)
				return
			}

			mux.ServeHTTP(writer, request)
			return
		}),
	)
}
