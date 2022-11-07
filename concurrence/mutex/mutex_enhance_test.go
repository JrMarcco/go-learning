package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	var mu Mutex
	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock() // 尝试获取到锁
	if ok {            // 获取成功
		fmt.Println("got the lock")
		// do something
		mu.Unlock()
		return
	}

	// 没有获取到
	fmt.Println("can't get the lock")
}

func TestCount(t *testing.T) {
	var mu Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()

			time.Sleep(time.Second)
		}()
	}

	time.Sleep(time.Second)
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t, starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

func TestDeadLock(t *testing.T) {
	var mu sync.RWMutex

	// writer,稍微等待，然后制造一个调用Lock的场景
	go func() {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("Lock")
		time.Sleep(100 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Unlock")
	}()

	go func() {
		factorial(&mu, 10) // 计算10的阶乘, 10!
	}()

	select {}
}

// 递归调用计算阶乘
func factorial(m *sync.RWMutex, n int) int {
	if n < 1 { // 阶乘退出条件
		return 0
	}
	fmt.Println("RLock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n // 递归调用
}
