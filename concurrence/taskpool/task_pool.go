package taskpool

type Task func()

type TaskPool struct {
	token chan struct{}
}

func New(limit int) *TaskPool {
	if limit <= 0 {
		limit = 8
	}
	p := &TaskPool{
		token: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		p.token <- struct{}{}
	}

	return p
}

func (p *TaskPool) Do(t Task) {
	<-p.token

	go func() {
		t()
		p.token <- struct{}{}
	}()
}

type CacheTaskPool struct {
	cache chan Task
}

func NewCachePool(limit int) *CacheTaskPool {
	if limit <= 0 {
		limit = 8
	}

	p := &CacheTaskPool{
		cache: make(chan Task, limit),
	}

	for i := 0; i < limit; i++ {
		go func() {
			for {
				select {
				case t, ok := <-p.cache:
					if !ok {
						return
					}
					t()
				}
			}
		}()
	}

	return p
}

func (p *CacheTaskPool) Do(t Task) {
	p.cache <- t
}
