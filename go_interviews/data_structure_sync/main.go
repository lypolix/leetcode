package main

import "sync"

type Stack struct{
	mutex sync.Mutex
	data []string
}

func NewStack() *Stack {
	return &Stack{}
}


func (b *Stack) Push(value string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.data = append(b.data, value)
}


func (b *Stack) Pop() {


	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.data) <= 0 {
		panic("pop: stack is empty")
	}

	b.data = b.data[:len(b.data) - 1]
}


func (b *Stack) Top() string {
	

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if len(b.data) <= 0 {
		panic("top: stack is empty")
	}
	
	return b.data[len(b.data) - 1]
}

var stack Stack


