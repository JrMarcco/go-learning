package mutex

// const (
// 	mutexLocked      = 1 << iota // 加锁标识位置
// 	mutexWoken                   // 唤醒标识位置
// 	mutexStarving                // 锁饥饿标识位置
// 	mutexWaiterShift = iota      // 标识 waiter 的起始 bit 位置
// )

// type Mutex struct {
// 	sync.Mutex
// }

// func (m *Mutex) TryLock() bool {

// 	// fastPath
// 	// 如果能成功抢到锁
// 	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
// 		return true
// 	}

// 	// 如果处于唤醒、加锁或者饥饿状态，则不参与竞争返回 false
// 	oldVal := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
// 	if oldVal&(mutexLocked|mutexStarving|mutexWoken) != 0 {
// 		return false
// 	}

// 	newVal := oldVal | mutexLocked
// 	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), oldVal, newVal)
// }

// func (m *Mutex) Count() int {
// 	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
// 	v = v>>mutexWaiterShift + (v & mutexLocked)
// 	return int(v)
// }

// func (m *Mutex) IsLocked() bool {
// 	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))&mutexLocked == mutexLocked
// }

// func (m *Mutex) IsWoken() bool {
// 	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))&mutexWoken == mutexWoken
// }

// func (m *Mutex) IsStarving() bool {
// 	return atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))&mutexStarving == mutexStarving
// }
