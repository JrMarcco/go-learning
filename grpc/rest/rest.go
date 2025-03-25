package rest

import (
	"context"
)

type ServiceImpl struct{}

func (s ServiceImpl) Get(_ context.Context, message *StringMessage) (*StringMessage, error) {
	return &StringMessage{
		Value: "Get hi: " + message.Value + "@",
	}, nil
}

func (s ServiceImpl) Post(_ context.Context, message *StringMessage) (*StringMessage, error) {
	return &StringMessage{
		Value: "Post hi: " + message.Value + "@",
	}, nil
}

func (s ServiceImpl) mustEmbedUnimplementedRestServiceServer() {
}
