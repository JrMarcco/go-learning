package distribute_lock

import (
	"sync"
	"testing"
	"time"
)

func TestIncr(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incr()
		}()

		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}
