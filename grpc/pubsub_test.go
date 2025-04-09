package grpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/moby/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPubsub(t *testing.T) {
	p := pubsub.NewPublisher(time.Second, 10)

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "golang:") {
				return true
			}
		}
		return false
	})

	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, "docker:") {
				return true
			}
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org/")
	go p.Publish("docker: https://www.docker.com/")

	go func() {
		fmt.Println("golang topic: ", <-golang)
	}()
	go func() {
		fmt.Println("docker topic: ", <-docker)
	}()
	<-make(chan bool)
}

func TestPubsubService_Server(t *testing.T) {
	grpcServer := grpc.NewServer()
	RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	_ = grpcServer.Serve(listener)
}

func TestPubsubService_Publish(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := NewPubsubServiceClient(conn)

	if _, err = client.Publish(
		context.Background(), &String{Value: "golang: hello Go"},
	); err != nil {
		log.Fatal(err)
	}

	if _, err = client.Publish(
		context.Background(), &String{Value: "docker: hello Docker"},
	); err != nil {
		log.Fatal(err)
	}
}

func TestPubsubService_Subscribe(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	client := NewPubsubServiceClient(conn)
	stream, err := client.Subscribe(
		context.Background(), &String{Value: "golang:"},
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if reply, err := stream.Recv(); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		} else {
			log.Println(reply.GetValue())
		}
	}
}
