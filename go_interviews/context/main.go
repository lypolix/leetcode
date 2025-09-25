package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main () {
	ctx, cansel := context.WithTimeout(context.Background(), time.Second)
	defer cansel()

	fmt.Println(callLongFunction(ctx))
}

func callLongFunction(ctx context.Context) (int, error) {
	result := make(chan int)

	go func() {
		defer close(result)
		result <- longFunction()
	}()

	for {
		select {
		case v := <- result: return v, nil
		case <-ctx.Done(): 
			return 0, errors.New("e")
		}
	}
}


func longFunction() int {
	time.Sleep(time.Duration(10 * time.Second))
	return 10
}
