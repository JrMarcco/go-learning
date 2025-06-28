package gpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	"github.com/JrMarcco/go-learning/gpc/message"
)

var _ Proxy = (*Server)(nil)

type Server struct {
	services map[string]ProxyStub
}

func (s *Server) Start(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		go s.handleConn(conn)
	}
}

func (s *Server) Call(ctx context.Context, req *message.Req) (*message.Resp, error) {
	resp := &message.Resp{}

	stub, ok := s.services[req.ServiceName]
	if !ok {
		return resp, fmt.Errorf("servicePtr %s not found", req.ServiceName)
	}

	respData, err := stub.Call(ctx, req.MethodName, req.Arg)
	if err != nil {
		return resp, err
	}

	resp.Data = respData
	return resp, nil
}

func (s *Server) handleConn(conn net.Conn) {
	for {
		bs, err := ReadMsg(conn)
		if err != nil {
			return
		}

		req := &message.Req{}
		err = json.Unmarshal(bs, req)
		if err != nil {
			return

		}
		resp, err := s.Call(context.Background(), req)
		if resp == nil {
			resp = &message.Resp{}
		}
		if err != nil && len(resp.Err) == 0 {
			resp.Err = err.Error()
		}

		encoded, err := s.encodeResp(resp)
		if err != nil {
			return
		}

		_, err = conn.Write(encoded)
		if err != nil {
			return
		}
	}
}

func (s *Server) encodeResp(src any) ([]byte, error) {
	respData, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	return EncodeMsg(respData), nil
}

func (s *Server) Register(service Service) {
	s.services[service.Name()] = &DefaultProxyStub{
		service: service,
		refVal:  reflect.ValueOf(service),
	}
}

func NewServer() *Server {
	return &Server{
		services: make(map[string]ProxyStub, 8),
	}
}
