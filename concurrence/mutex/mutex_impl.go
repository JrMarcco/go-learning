package mutex

import (
	"sync/atomic"
)

// mutex 实现的 4 个阶段

// cas
// CAS 指令，当时还没有抽象出 atomic 包
// 将给定的值和一个内存地址中的值进行比较，
// 如果相等则使用新的值替换地址中的值。
// 这个操作是原子性的。
func cas(val *int32, old, new int32) bool { return true }
func semaAcquire(sema *int32)             { return }
func semaRelease(sema *int32)             { return }

// MutexV1 第一阶段。
// 互斥锁的结构包含两个字段。
//
// 注意：
//
//	MutexV1 本身没有包含持有这把锁的 goroutine 信息。
type MutexV1 struct {
	key  int32 // 锁是否被持有的标识
	sema int32 // 信号量，用以 阻塞/唤醒 goroutine
}

// xadd 保证成功再 val 上增加 delta 的值。
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) {
			return v + delta
		}
	}

}

// Lock 请求锁
func (m *MutexV1) Lock() {
	if xadd(&m.key, 1) == 1 {
		// 标识 + 1，如果结果为 1，则表示成功获取锁
		return
	}
	// 阻塞等待
	// 阻塞等待的 goroutine 会进入信号量的等待队列
	// 即排队获取锁
	semaAcquire(&m.sema)
}

// Unlock 释放锁
//
// 注意：
//
//	Unlock 可以被任意 goroutine 调用，即使这个 goroutine 没有持有锁。
//	MutexV1 本身没有包含持有这把锁的 goroutine 信息，所以 Unlock 也无法对此进行检查。
//	** 这就导致其他 goroutine 可以调用 Unlock 释放一把未被持有的锁 **。
func (m *MutexV1) Unlock() {
	if xadd(&m.key, -1) == 0 {
		// 释放锁
		// key 为 0 表示没有其他等待的 goroutine
		return
	}
	// 唤醒等待的 goroutine
	semaRelease(&m.sema)
}

const (
	mutexLocked      = 1 << iota // 1 << 0 = 1 （ 二进制：00000001 ）
	mutexWoken                   // 1 << 1 = 2 （ 二进制：00000010 ）
	mutexWaiterShift = iota      // 此时 iota = 2
)

// MutexV2 第二阶段。
type MutexV2 struct {
	// state 是一个 int32 类型的变量
	// 布局如下：
	//
	// 	|   31...3   |      2      |     1      |      0     |
	// 	|   waiters  |    waiter   |   woken    |   locked   |
	// 	|   count    |    shift    |   flag     |   status   |
	state int32
	sema  uint32
}

func (m *MutexV2) Lock() {
	// Fast path: 快速路径，直接尝试获取锁
	// 在运气好的情况下能直接获取到锁
	// 这样能让 CPU 中正在执行的 goroutine 有更多机会获取到锁
	// 一定程度上提高性能
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	// slow path: 慢速路径
	awoke := false
	// 不断尝试获取锁
	for {
		old := m.state

		new := old | mutexLocked

		// mutexLocked = 1 （二进制：00000001）
		// old 为当前锁的状态，
		// old & mutexLocked != 0
		// |-- 结果为 true，表示 m.state 的最低位为 1，即锁被持有
		// |-- 结果为 false，表示 m.state 的最低位为 0，即锁未被持有
		//
		// 假设：
		// old = 5 ( 二进制：00000101 )
		//
		// 	old & mutexLocked
		// ->
		//   00000101
		// & 00000001
		// 	----------
		//   00000001
		// 即 old 最低位为 1，锁已经被持有
		if old&mutexLocked != 0 {
			// waiters + 1
			//
			// mutexWaiterShift = 2
			// 1<<mutexWaiterShift = 1<<2 = 4 （ 二进制：00000100 ）
			// waiters count 存储在 state 的 3...31 位
			// 所以需要将 1 左移 2 位，即 4 （ 二进制：00000100 ）
			// 即每次 +4 相当于增加一个等待者
			new = old + 1<<mutexWaiterShift
		}

		if awoke {
			// &^ （ AND NOT ）按位清除运算
			// 将 new 的 mutexWoken 位置 0
			// 即 goroutine 被唤醒
			// 新状态清除唤醒标记
			new &^= mutexWoken
		}

		// 尝试更新状态
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&mutexLocked == 0 {
				// 锁原来处于未加锁状态（ old 最低位为 0 ）
				// 获取锁成功
				break
			}

			// 请求信号量，阻塞等待
			// runtime_Semacquire 是 runtime 包中的私有方法
			// runtime_Semacquire(&m.sema)
			awoke = true
		}
	}
}
