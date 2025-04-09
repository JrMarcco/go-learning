package v4

import (
	"sync/atomic"
	"time"
)

const (
	mutexLocked      = 1 << iota // 1 << 0 = 1 （ 二进制：00000001 ）
	mutexWoken                   // 1 << 1 = 2 （ 二进制：00000010 ）
	mutexStarving                // 1 << 2 = 4 （ 二进制：00000100 ）
	mutexWaiterShift = iota      // 此时 iota = 3

	starvingThresholdNs = 1e6 // 1 毫秒
)

// MutexV4 第四阶段
type MutexV4 struct {
	// state 是一个 int32 类型的变量
	// 布局如下：
	//
	// 	|   31...3   |     2      |     1      |      0     |
	// 	|   waiters  |   starving |   woken    |   locked   |
	// 	|   count    |   flag     |    flag    |   status   |
	state int32
	sema  uint32
}

func (m *MutexV4) Lock() {
	// Fast path: 快速路径，直接尝试获取锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	m.lockSlow()
}

// lockSlow
// Slow path: 慢速路径
func (m *MutexV4) lockSlow() {
	var waitStartTime int64
	// 当前 goroutine 的饥饿标记
	starving := false
	awoke := false
	iter := 0
	old := m.state

	for {
		// mutexLocked|mutexStarving => 1 | 4 = 5 （ 二进制：00000101 ）
		// old&(mutexLocked|mutexStarving) == mutexLocked
		// 表示锁被持有，即最低位为 1
		// 并且锁没有处于饥饿状态（ 锁的 mutexStarving 位为 0 ）
		// 并且可以尝试自旋
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// 如果当前 goroutine 没有被唤醒
			// 并且锁没有被唤醒（ 锁的 mutexWoken 位为 0 ）
			// 并且等待者数量不为 0
			// 则尝试唤醒锁（ old|mutexWoken， 即设置锁的 mutexWoken 位为 1 ）
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 && atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}

			// 尝试自旋
			runtime_doSpin()
			iter++
			continue
		}

		new := old
		// 如果锁没有处于饥饿状态，则尝试加锁
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}

		// 如果锁被持有或者处于饥饿状态，则增加等待者数量
		if old&(mutexLocked|mutexStarving) != 0 {
			// waiters + 1
			new += 1 << mutexWaiterShift
		}

		// 如果当前 goroutine 处于饥饿状态
		// 并且锁被持有
		// 则设置饥饿标记
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}

		// awoke 为 true 表示当前 goroutine 被唤醒
		if awoke {
			// 状态一致性检查
			if new&mutexWoken == 0 {
				panic("sync: inconsistent mutex state")
			}
			// 新状态清除唤醒标记
			new &^= mutexWoken
		}

		// 成功设置新的锁状态
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			// 如果锁原来处于未加锁状态
			// 并且锁没有处于饥饿状态
			// 获取锁并返回
			if old&(mutexLocked|mutexStarving) == 0 {
				break
			}

			// 处理饥饿状态
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}

			// 阻塞等待
			runtime_Semacquire(&m.sema, queueLifo, 1)

			// 唤醒后检查锁是否应该处于饥饿状态
			// runtime_nanotime()-waitStartTime > starvingThresholdNs
			// 表示当前 goroutine 等待锁的时间超过了阈值 （ 1 毫秒 ）
			starving = starving || runtime_nanotime()-waitStartTime > starvingThresholdNs

			old = m.state

			// old&mutexStarving != 0 为 true 表示锁处于饥饿状态
			if old&mutexStarving != 0 {
				// 状态一致性检查
				// 饥饿模式的引入是为了防止某些协程长时间获取不到锁（ 饿死 ）
				//
				// 所以当处于饥饿模式下时
				// 1. mutexLocked 必须为 0，即 锁已经被释放
				//		锁在被持有的情况下不会进入饥饿模式的转移逻辑
				// 2. mutexWoken 必须为 0，即 锁没有被唤醒
				//		这是因为饥饿模式下禁止通过唤醒（ woken ）机制抢占锁
				//		必须严格按照队列顺序传递锁
				// 3. old>>mutexWaiterShift == 0，即 waiters == 0
				//		饥饿模式解决的是长时间等待的问题，没有等待者说明没有协程需要锁
				//		与饥饿模式相悖
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					panic("sync: inconsistent mutex state")
				}

				// 加锁
				// 并且 waiters - 1
				delta := int32(mutexLocked - 1<<mutexWaiterShift)

				// 如果锁不处于饥饿状态
				// 或者等待者数量为 1（ 即最后一个 waiter ）
				if !starving || old>>mutexWaiterShift == 1 {
					// 清除饥饿标记
					delta -= mutexStarving
				}

				atomic.AddInt32(&m.state, delta)
				break
			}

			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}
}

func (m *MutexV4) Unlock() {
	// Fast path: 快速路径，直接尝试释放锁
	// Drop lock bit，即去掉锁标记
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		m.unlockSlow(new)
	}
}

func (m *MutexV4) unlockSlow(new int32) {
	if (new+mutexLocked)&mutexLocked == 0 {
		panic("sync: unlock of unlocked mutex")
	}

	if new&mutexStarving == 0 {
		old := new
		for {
			// 如果等待者数量为 0
			// 或者锁被持有
			// 或者锁被唤醒
			// 或者锁处于饥饿状态（ 饥饿模式下锁必须严格按照 FIFO 顺序传递，不能随意唤醒协程，否则会破坏公平性 ）
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}

			// waiters - 1
			// 设置唤醒标记
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, true, 1)
				return
			}

			old = m.state
		}
	} else {
		runtime_Semrelease(&m.sema, true, 1)
	}
}

func runtime_Semacquire(sema *uint32, lifo bool, delta int32) {}
func runtime_Semrelease(sema *uint32, handoff bool, nr int32) {}
func runtime_canSpin(iter int) bool {
	return true
}
func runtime_doSpin() {}
func runtime_nanotime() int64 {
	return time.Now().UnixNano()
}
