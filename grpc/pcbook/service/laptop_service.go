package service

import (
	"context"
	"github.com/google/uuid"
	"go-learning/grpc/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type LaptopServer struct {
	laptopStore LaptopStore
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

func (l *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopReq) (*pb.CreateLaptopRsp, error) {

	laptop := req.GetLaptop()
	log.Printf("receive a create laptop req with id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		if _, err := uuid.Parse(laptop.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid: %v", err)
		}
	} else {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop id: %v", err)
		}
		laptop.Id = uuid.String()
	}
	return nil, nil
}
