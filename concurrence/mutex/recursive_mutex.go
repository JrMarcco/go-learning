package mutex

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

func (rm *RecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&rm.owner) == token {
		rm.recursion++
		return
	}

	rm.Mutex.Lock()
	atomic.StoreInt64(&rm.owner, token)
	rm.recursion = 1
}

func (rm *RecursiveMutex) UnLock(token int64) {
	if atomic.LoadInt64(&rm.owner) != token {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", rm.owner, token))
	}
	rm.recursion--
	if rm.recursion != 0 {
		return
	}
	atomic.StoreInt64(&rm.owner, -1)
	rm.Mutex.Unlock()
}
