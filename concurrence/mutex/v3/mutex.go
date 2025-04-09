package v3

import "sync/atomic"

const (
	mutexLocked      = 1 << iota // 1 << 0 = 1 （ 二进制：00000001 ）
	mutexWoken                   // 1 << 1 = 2 （ 二进制：00000010 ）
	mutexWaiterShift = iota      // 此时 iota = 2
)

// MutexV3 第三阶段
type MutexV3 struct {
	// state 是一个 int32 类型的变量
	// 布局如下：
	//
	// 	|   31...2   |     1      |      0     |
	// 	|   waiters  |   woken    |   locked   |
	// 	|   count    |   flag     |   status   |
	state int32
	sema  uint32
}

// Lock 请求锁
// 在 MutexV2 的基础上，增加了自旋（ spin ）
func (m *MutexV3) Lock() {
	// Fast path: 快速路径，直接尝试获取锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	// Slow path: 慢速路径
	awoke := false
	iter := 0

	for {
		old := m.state
		new := old | mutexLocked

		// old&mutexLocked != 0 为 true 表示锁被持有
		if old&mutexLocked != 0 {
			// 如果可以自旋，则尝试自旋
			if runtime_canSpin(iter) {
				// 如果当前 goroutine 没有被唤醒
				// 并且锁没有被唤醒
				// 并且等待者数量不为 0
				// 则尝试唤醒锁
				if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 && atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
					awoke = true
				}

				// 自旋
				// 自旋的目的是为了让当前 goroutine 有更多机会获取到锁
				// 即不休眠竞争锁
				// 减少上下文切换
				runtime_doSpin()
				iter++
				continue
			}

			// waiters + 1
			new = old + 1<<mutexWaiterShift
		}

		// awoke 为 true 表示当前 goroutine 被唤醒
		if awoke {
			// 状态一致性检查
			// awoke 为 true 表示当前 goroutine 被唤醒
			// 那么 new 的唤醒标记应该为 1
			// 即锁的 mutexWoken 位应该被设置为 1
			if new&mutexWoken == 0 {
				panic("sync: inconsistent mutex state")
			}

			// 新状态清除唤醒标记
			new &^= mutexWoken
		}

		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			// 如果锁原来处于未加锁状态
			// 则获取锁成功
			if old&mutexLocked == 0 {
				break
			}

			// 请求信号量，阻塞等待
			runtime_Semacquire(&m.sema)
			awoke = true
			iter = 0
		}
	}
}

// Unlock 释放锁
// 同 MutexV2
func (m *MutexV3) Unlock() {
	// Fast path: 快速路径，直接尝试释放锁
	// Drop lock bit，即去掉锁标记
	new := atomic.AddInt32(&m.state, -mutexLocked)

	// (new+mutexLocked)&mutexLocked == 0
	//
	// 	(new+mutexLocked) 为 new 加上 mutexLocked 的值，即恢复为原来的状态 old
	// 	(new+mutexLocked)&mutexLocked == 0
	// 	相当于 old & mutexLocked == 0
	//  即 old 最低位为 0，表示为加锁状态
	if (new+mutexLocked)&mutexLocked == 0 {
		panic("sync: unlock of unlocked mutex")
	}

	// Slow path: 慢速路径
	old := new
	for {
		// old>>mutexWaiterShift 为当前等待者的数量（waiters count 存储在 state 的 2...31 位）
		// state 右移 2 位，即为当前等待者的数量
		//
		// mutexLocked|mutexWoken => 1 | 2 = 3 （ 二进制：00000011 ）
		// 最低位为 1，表示锁被持有（ 检查期间其他 goroutine 获取到锁 ）
		// 次低位为 1，表示锁被唤醒（ 检查期间其他 goroutine 被唤醒 ）
		// 所以 old&(mutexLocked|mutexWoken) != 0 表示锁被持有或者被唤醒
		if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 {
			return
		}

		// 1<<mutexWaiterShift 表示 1 个等待者在 state 中的位置（waiters count 存储在 state 的 2...31 位）
		// old - 1<<mutexWaiterShift 表示当前等待者的数量减 1
		// | mutexWoken 确保 new 的唤醒标记成功设置
		new = (old - 1<<mutexWaiterShift) | mutexWoken
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			runtime_Semrelease(&m.sema)
			return
		}

		// 如果 CompareAndSwapInt32 失败，则更新 old 的值并重试
		// 这里是为了处理并发竞争
		// CAS 阶段可能有其他的 goroutine 修改了 state 的值
		// 所以需要重新读取 state 的值来计算 new 状态
		old = m.state
	}
}

func runtime_Semacquire(sema *uint32) {}
func runtime_Semrelease(sema *uint32) {}
func runtime_canSpin(iter int) bool {
	return true
}
func runtime_doSpin() {}
