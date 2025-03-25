package sema

import (
	"context"
	"golang.org/x/sync/semaphore"
	"runtime"
	"testing"
	"time"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers))
	task       = make([]int, 4*maxWorkers)
)

func TestSema(t *testing.T) {
	ctx := context.Background()

	for i := range task {
		if err := sema.Acquire(ctx, 1); err != nil {
			t.Log(err)
			break
		}

		go func(i int) {
			defer sema.Release(1)

			time.Sleep(time.Second)
			task[i] = i + 1
		}(i)
	}

	// 在实际应用中如果想等待所有的 worker 都执行完，
	// 就可以获取最大计数值的信号量。
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		t.Log(err)
	}

	t.Log(task)
}
