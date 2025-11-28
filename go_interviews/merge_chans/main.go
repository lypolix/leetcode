package main

import (
	"fmt"
	"sync"
)



func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go func() {
		for i := 0; i < 10; i ++ {
			ch1 <- struct{}{}
		}
		close(ch1)
	}()
	go func() {
		for i := 0; i < 10; i ++ {
			ch2 <- struct{}{}
		}
		close(ch2)
	}()
	go func() {
		for i := 0; i < 10; i ++ {
			ch3 <- struct{}{}
		}
		close(ch3)
	}()

	res := merge(ch1, ch2, ch3)
	for i := range res {
		fmt.Println(i)
	}
}


func merge(chans ... chan struct{}) <-chan struct{}{
	result := make(chan struct{})
	wg := sync.WaitGroup{}
	for _, ch := range chans {
		wg.Add(1)
		go func(ch chan struct{}) {
			defer wg.Done()
			for i := range ch {
				result <- i
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}


