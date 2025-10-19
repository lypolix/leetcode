package main

import (

	"sync/atomic"
	"unsafe"
)

type item struct {
	value int
	next unsafe.Pointer
}

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewQueue() Queue {
	dummy := &item{}
	return Queue{
		head: unsafe.Pointer(dummy),
		tail: unsafe.Pointer(dummy),
	}
}

func (q *Queue) Push(value int) {
	node := &item{value: value}

	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*item)(tail).next)


		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				if atomic.CompareAndSwapPointer(&(*item)(tail).next, next, unsafe.Pointer(node)) {
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
					return
				}
			}else {
				atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(node))
			}
		}
	}
}


func (q *Queue) Pop() int {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*item)(head).next)

		if head == atomic.LoadPointer(&q.head) {
			if head == tail {
				if next == nil {
					return -1
				}else {
					atomic.CompareAndSwapPointer(&q.tail, tail, next)
				}

			} else {
				value := (*item)(next).value
				if atomic.CompareAndSwapPointer(&q.head, head, next){
					return value
				}
			}
		}

	}
}