package service

import (
	"context"
	"github.com/google/uuid"
	"go-learning/grpc/pcbook/db"
	"go-learning/grpc/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type LaptopService struct {
	client *db.Client
}

func NewLaptopServer() *LaptopService {
	return &LaptopService{
		client: entClient,
	}
}

func (l *LaptopService) CreateLaptop(ctx context.Context, req *pb.CreateLaptopReq) (*pb.CreateLaptopRsp, error) {

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

	lp, err := l.client.Laptop.Create().
		SetUID(laptop.Id).
		SetBrand(laptop.GetBrand()).
		SetName(laptop.GetBrand()).
		Save(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot save laptop: %v", err)
	}

	return &pb.CreateLaptopRsp{Id: lp.ID, Uuid: lp.UID}, nil
}
