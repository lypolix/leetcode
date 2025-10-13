package main

import (
	"context"
	"fmt"
	"sync"
)

func Start(ctx context.Context, workersCount int, input <-chan int, transform func(e int) int) <-chan int {
	result := make(chan int)
	wg := new(sync.WaitGroup)

	for i := 0; i < workersCount; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return 
				case v, ok := <-input:
					if !ok {
						return 
					}

					select {
					case <- ctx.Done():
						return 
					case result <- transform(v):
					}



				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func main(){

	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int, 10)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	result := Start(ctx, 1, ch, func(e int) int {
		return e + 1
	})

	cancel()
	for e := range result {
		fmt.Println(e)
	}
}