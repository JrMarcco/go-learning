package gpc

import (
	"context"

	"github.com/JrMarcco/go-learning/gpc/message"
)

type Proxy interface {
	Call(ctx context.Context, req *message.Req) (*message.Resp, error)
}

type ProxyStub interface {
	Call(ctx context.Context, methodName string, arg []byte) ([]byte, error)
}

type Service interface {
	Name() string
}
