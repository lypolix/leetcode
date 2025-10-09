package main

import (
	"fmt"
	"time"
)

func run() {
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	fmt.Println("result:", <-ch)
	time.Sleep(2 * time.Second)
}

func main() {
	run()
}
