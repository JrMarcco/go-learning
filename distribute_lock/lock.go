package distribute_lock

type Lock struct {
	ch chan struct{}
}

func NewLock() Lock {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	return Lock{
		ch: ch,
	}
}

func (l Lock) Lock() bool {
	select {
	case <-l.ch:
		return true
	default:
		return false
	}
}

func (l Lock) Unlock() {
	l.ch <- struct{}{}
}
