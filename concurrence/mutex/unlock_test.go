package mutex

import (
	"fmt"
	"sync"
	"testing"
)

/*
*
10个goroutine同时对count进行100w次+1操作（不加锁）
*/
func TestUnlock(t *testing.T) {
	var count = 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

// 如果嵌入的 struct 有多个字段，我们一般会把 Mutex 放在要控制的字段上面，然后使用空格把字段分隔开来
type counter struct {
	mu    sync.Mutex
	count uint64
}

func TestCounter(t *testing.T) {
	var counter counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 1000000; j++ {
				counter.mu.Lock()
				counter.count++
				counter.mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}

type richCounter struct {
	counterType int
	name        string

	mu    sync.Mutex
	count uint64
}

func (r *richCounter) Incr() {
	r.mu.Lock()
	r.count++
	r.mu.Unlock()
}

func (r *richCounter) Count() uint64 {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.count
}

func TestRichCounter(t *testing.T) {
	var richCounter richCounter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				richCounter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(richCounter.Count())

}
