package distribute_lock

import (
	"sync"
	"testing"
)

func TestLock(t *testing.T) {

	counter := 0
	l := NewLock()
	var wg = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if !l.Lock() {
				t.Log("fail to try lock")
				return
			}

			counter++
			t.Log("current counter: ", counter)
			l.Unlock()
		}()
	}

	wg.Wait()
}
