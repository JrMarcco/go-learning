package busi_alg

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶
type TokenBucket struct {
	rate   int64 // 令牌放入速度
	max    int64 // 令牌最大数量
	last   int64 // 上一次请求发生时间
	amount int64 // 令牌数量
	m      sync.Mutex
}

func cur() int64 {
	return time.Now().Unix()
}

func New(rate, max int64) *TokenBucket {
	// TODO check max & rate
	return &TokenBucket{
		rate:   rate,
		max:    max,
		amount: max,
		last:   cur(),
	}
}

func (t *TokenBucket) Pass() bool {
	t.m.Lock()
	defer t.m.Unlock()

	// 时间间隔
	interval := cur() - t.last

	// 超过最大数量则不继续补充令牌
	amount := t.amount + interval*t.rate
	if amount > t.max {
		amount = t.max
	}

	// 没有令牌则请求不通过
	if amount <= 0 {
		return false
	}

	// 取出令牌
	amount--
	t.amount = amount
	t.last = cur()

	return true
}
