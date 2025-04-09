package service

import (
	"context"

	"log"

	"github.com/JrMarcco/go-learning/grpc/pcbook/db"
	"github.com/JrMarcco/go-learning/grpc/pcbook/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	log.Printf("receive a create laptop req with id: %s", laptop.Uid)
	if len(laptop.Uid) > 0 {
		if _, err := uuid.Parse(laptop.Uid); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop id is not a valid uuid: %v", err)
		}
	} else {
		uid, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop id: %v", err)
		}
		laptop.Uid = uid.String()
	}

	lp, err := l.client.Laptop.Create().
		SetUID(laptop.Uid).
		SetBrand(laptop.GetBrand()).
		SetLaptopName(laptop.GetBrand()).
		SetWeight(2.25).
		SetPriceRmb(6999).
		SetReleaseYear(2010).
		Save(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot save laptop: %v", err)
	}

	return &pb.CreateLaptopRsp{Id: lp.ID, Uid: lp.UID}, nil
}
