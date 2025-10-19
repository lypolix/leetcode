package main

import (
	"sync/atomic"
	"unsafe"
)


type item struct {
	value int
	next unsafe.Pointer
}

type Stack struct {
	head unsafe.Pointer
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	node := &item{value: value, next: atomic.LoadPointer(&s.head)}
	atomic.StorePointer(&s.head, unsafe.Pointer(node))
}

func (s *Stack) Pop() int {
	head := atomic.LoadPointer(&s.head)
	if head == nil {
		return -1
	}
	value := (*item)(head).value
	next := atomic.LoadPointer(&(*item)(head).next)
	atomic.StorePointer(&s.head, next)
	return value
}