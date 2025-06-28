package cond

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})

	ready := 0

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()

			t.Logf("%d ready", i)
			c.Broadcast()
		}(i)
	}

	// 调用 Wait() 之前要先加锁，这是为了保证条件判断的原子性。
	// Wait() 内部会自动解锁 -> 休眠 -> 唤醒后重新加锁。
	// 如果没有预先加锁会导致解锁时候 panic，甚至出现竞态条件，
	// 锁保护的数据也无法安全访问。
	// 官方文档明确要求：
	//		   	c.L.Lock()
	//   		for !condition() {
	//       		c.Wait()
	//   		}
	//   		c.L.Unlock()
	c.L.Lock()
	for ready != 10 {
		c.Wait()
		t.Log("waiter wake up once")
	}
	c.L.Unlock()

	t.Log("all ready")
}
