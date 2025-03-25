package taskpool

import (
	"testing"
	"time"
)

func TestTaskPool(t *testing.T) {
	//p := New(8)
	p := NewCachePool(8)

	for i := 0; i < 10; i++ {
		p.Do(func() {
			t.Log("doing")
			time.Sleep(time.Second)
			t.Log("done")
		})
	}

	time.Sleep(2 * time.Second)
	t.Log("all done")
}
