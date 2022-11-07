package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"testing"
)

func TestService(t *testing.T) {
	grpcServer := grpc.NewServer()
	RegisterRestServiceServer(grpcServer, new(ServiceImpl))

	lis, _ := net.Listen("tcp", ":8081")
	_ = grpcServer.Serve(lis)
}

func TestRestGw(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	err := RegisterRestServiceHandlerFromEndpoint(
		ctx, mux, ":8081",
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	_ = http.ListenAndServe(":8080", mux)
}
