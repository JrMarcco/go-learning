package once

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	sync.Mutex
	done uint32
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.doSlow(f)
}

func (o *Once) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}

func (o *Once) doSlow(f func() error) error {
	o.Lock()
	defer o.Unlock()

	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			// 执行成功再设置标记
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return nil
}
