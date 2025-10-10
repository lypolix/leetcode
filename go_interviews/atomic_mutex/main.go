package main

import "sync/atomic"


type Mutex struct {
	state int32
}


func (m *Mutex) Lock() {
	
	for !atomic.CompareAndSwapInt32(&m.state, 0, 1) {

	}
}

func (m *Mutex) Unlock() {
	atomic.StoreInt32(&m.state, 0)
}

func main() {
	
}