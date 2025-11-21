package main

import (
	"bufio"
	"fmt"
	"os"
)



type Heap struct {
	data []int
}


func (h *Heap) Insert(x int) {
	h.data = append(h.data, x)
	h.liftUp(len(h.data) - 1)
}

func (h *Heap) liftUp(i int) {
	for i > 0 {
		p := (i - 1)/2
		if h.data[p] >= h.data[i] {
			break
		}
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

func (h *Heap) Extract() int {
	res := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	h.liftDown(0)
	return res
}

func (h *Heap) liftDown(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i
		if left < n && h.data[left] > h.data[largest] {
            largest = left
        }
        if right < n && h.data[right] > h.data[largest] {
            largest = right
        }
        if largest == i {
            break
        }
        h.data[i], h.data[largest] = h.data[largest], h.data[i]
        i = largest
	}	
}


func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	heap := &Heap{}
	var a, b int
	for i := 0; i < n; i ++ {
		fmt.Fscan(in, &a)
		if a == 0 {
			fmt.Fscan(in, &b)
			heap.Insert(b)
		} else {
			fmt.Println(heap.Extract())
		}
	}
}
	