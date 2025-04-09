package v1

// MutexV1 第一阶段
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

func cas(val *int32, old, new int32) bool { return true }
func semaAcquire(sema *int32)             { return }
func semaRelease(sema *int32)             { return }
