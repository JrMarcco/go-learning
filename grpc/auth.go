package grpc

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authentication struct {
	Username string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"username": a.Username,
		"password": a.Password,
	}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	// 这里为了简单返回 false 表示不要求底层使用安全链接
	return false
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("missing credentials")
	}

	var username, password string
	if val, ok := md["username"]; ok {
		username = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}

	if username == a.Username && password == a.Password {
		return nil
	}
	return status.Errorf(codes.Unauthenticated, "invalid token")
}
