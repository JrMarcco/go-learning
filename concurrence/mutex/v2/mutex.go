package v2

import "sync/atomic"

const (
	mutexLocked      = 1 << iota // 1 << 0 = 1 （ 二进制：00000001 ）
	mutexWoken                   // 1 << 1 = 2 （ 二进制：00000010 ）
	mutexWaiterShift = iota      // 此时 iota = 2
)

// MutexV2 第二阶段
type MutexV2 struct {
	// state 是一个 int32 类型的变量
	// 布局如下：
	//
	// 	|   31...2   |     1      |      0     |
	// 	|   waiters  |   woken    |   locked   |
	// 	|   count    |   flag     |   status   |
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
			// waiters count 存储在 state 的 2...31 位
			// 所以需要将 1 左移 2 位，即 4 （ 二进制：00000100 ）
			// 即每次 +4 相当于增加一个等待者
			new = old + 1<<mutexWaiterShift
		}

		// awoke 为 true 表示当前 goroutine 被唤醒
		if awoke {
			// &^ （ AND NOT ）按位清除运算
			// 将 new 的 mutexWoken 位置 0
			// 即 goroutine 被唤醒
			// 新状态清除唤醒标记
			// 即 new 的 mutexWoken 位设置为 0
			new &^= mutexWoken
		}

		// 尝试更新状态
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			// 即使成功设置 state 的值
			// 也不能保证锁被成功获取
			// 如果之前的 state 值是锁被持有的状态
			// 那么 state 只会清除 mutexWoken 标记或者增加一个 waiter
			// 不会改变锁的状态
			if old&mutexLocked == 0 {
				// 锁原来处于未加锁状态（ old 最低位为 0 ）
				// 获取锁成功
				break
			}

			// 请求信号量，阻塞等待
			runtime_Semacquire(&m.sema)
			awoke = true
		}
	}
}

func (m *MutexV2) Unlock() {
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
		// 最低位为 1，表示锁被持有
		// 次低位为 1，表示锁被唤醒
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
