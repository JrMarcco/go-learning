package gpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type MockReq struct {
	Name string
}

type MockResp struct {
	Msg string
}

var _ Service = (*MockClientService)(nil)

type MockClientService struct {
	SayHello func(ctx context.Context, req *MockReq) (*MockResp, error)
}

func (m *MockClientService) Name() string {
	return "mock-servicePtr"
}

var _ Service = (*MockServerService)(nil)

type MockServerService struct {
}

func (m *MockServerService) Name() string {
	return "mock-servicePtr"
}

func (m *MockServerService) SayHello(_ context.Context, req *MockReq) (*MockResp, error) {
	return &MockResp{
		Msg: fmt.Sprintf("hello %s", req.Name),
	}, nil
}

func TestClientProxy(t *testing.T) {
	svr := NewServer()
	serverSvc := &MockServerService{}
	svr.Register(serverSvc)

	go func() {
		err := svr.Start(":8081")
		require.NoError(t, err)
	}()
	time.Sleep(time.Second)

	clientSvc := &MockClientService{}
	client, err := NewClient(":8081")
	require.NoError(t, err)

	setProxyFunc(clientSvc, client)

	resp, err := clientSvc.SayHello(context.Background(), &MockReq{Name: "jrmarcco"})
	require.NoError(t, err)
	require.Equal(t, "hello jrmarcco", resp.Msg)
}
