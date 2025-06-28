package gpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/JrMarcco/go-learning/gpc/message"
	"github.com/silenceper/pool"
)

var _ Proxy = (*Client)(nil)

type Client struct {
	pool pool.Pool
}

func (c *Client) Call(_ context.Context, req *message.Req) (*message.Resp, error) {
	val, err := c.pool.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to get available connection: %w", err)
	}

	defer c.pool.Put(val)

	conn := val.(net.Conn)

	bs, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to json marshal req: %w", err)
	}

	encoded := EncodeMsg(bs)
	_, err = conn.Write(encoded)
	if err != nil {
		return nil, err
	}

	bs, err = ReadMsg(conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read msg: %w", err)
	}

	resp := &message.Resp{}
	err = json.Unmarshal(bs, resp)
	return resp, err

}

func NewClient(addr string) (*Client, error) {
	p, err := pool.NewChannelPool(&pool.Config{
		InitialCap:  8,
		MaxCap:      64,
		MaxIdle:     16,
		IdleTimeout: time.Minute,
		Factory: func() (any, error) {
			return net.Dial("tcp", addr)
		},
		Close: func(v any) error {
			return v.(net.Conn).Close()
		},
	})
	if err != nil {
		return nil, err
	}
	return &Client{pool: p}, nil
}
