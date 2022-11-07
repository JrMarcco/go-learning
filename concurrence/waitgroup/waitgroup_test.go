package waitgroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Counter struct {
	sync.RWMutex
	count uint64
}

func (c *Counter) incr() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *Counter) Count() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

func TestCounter(t *testing.T) {
	counter := &Counter{count: 10}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Second)
			counter.incr()
		}()
	}

	wg.Wait()

	fmt.Println(counter.count)
}
