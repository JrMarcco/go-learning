package grpc

import (
	"context"
	"io"
	"log"
)

type HelloServiceImpl struct {
	auth *Authentication
}

func NewHelloServiceImpl() *HelloServiceImpl {
	return &HelloServiceImpl{
		auth: &Authentication{
			Username: "jrmarcco",
			Password: "util",
		},
	}
}

func (h *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {
	panic("implement me")
}

func (h *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	if err := h.auth.Auth(ctx); err != nil {
		return nil, err
	}
	return &String{
		Value: "hello: " + args.GetValue(),
	}, nil
}

// Channel 使用 grpc 流双向通信
// 双向流数据的发送和接受都是完全独立的行为
// 发送和接受的操作不需要一一对应
func (h *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return nil
		}

		log.Printf("receive msg from client: %s\n", args.GetValue())

		if err = stream.Send(&String{Value: "hey"}); err != nil {
			return err
		}
	}
}
