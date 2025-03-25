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

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		t.Log("waiter wake up once")
	}
	c.L.Unlock()

	t.Log("all ready")
}
