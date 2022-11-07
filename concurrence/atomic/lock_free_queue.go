package atomic

import (
	"sync/atomic"
	"unsafe"
)

type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node struct {
	val  any
	next unsafe.Pointer
}

func NewLKQueue() *LKQueue {
	pn := unsafe.Pointer(&node{})
	return &LKQueue{
		head: pn,
		tail: pn,
	}
}

func (q *LKQueue) Enqueue(val any) {

	n := &node{val: val}

	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == load(&q.tail) {
			if next == nil {
				if cas(&tail.next, next, n) {
					cas(&q.tail, tail, n)
					return
				}
			} else {
				cas(&q.tail, tail, next)
			}
		}
	}
}

func (q *LKQueue) Dequeue() any {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)

		if head == load(&q.head) {
			if head == tail {
				if next == nil {
					return nil
				}

				cas(&q.tail, tail, next)
			} else {
				v := next.val
				if cas(&q.head, head, next) {
					return v
				}
			}
		}
	}
}

func load(p *unsafe.Pointer) *node {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new),
	)
}
