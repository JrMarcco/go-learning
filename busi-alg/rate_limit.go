package busi_alg

import (
	"sync"
	"time"
)

type RateLimiter struct {
	rate   int64 // 令牌放入速度
	max    int64 // 令牌最大数量
	last   int64 // 上一次请求发生时间
	amount int64 // 令牌数量
	m      sync.Mutex
}

func cur() int64 {
	return time.Now().Unix()
}

func New(rate, max int64) *RateLimiter {
	// TODO check max & rate
	return &RateLimiter{
		rate:   rate,
		max:    max,
		amount: max,
		last:   cur(),
	}
}

func (r *RateLimiter) Pass() bool {
	r.m.Lock()
	defer r.m.Unlock()

	// 时间间隔
	interval := cur() - r.last

	amount := r.amount + interval*r.rate
	if amount > r.max {
		amount = r.max
	}

	amount--
	r.amount = amount
	r.last = cur()
	return true
}
