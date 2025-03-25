package mutex

import (
	"fmt"
	"sync"
	"testing"
)

/*
*
10个goroutine同时对count进行100w次+1操作（加锁）
*/
func TestLock(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				count++
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
