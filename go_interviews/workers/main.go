package main

import (
	"fmt"
	"time"
)


func main() {
	timeStart := time.Now()
	_, _ = <-worker(), <-worker()
	fmt.Println(int(time.Since(timeStart).Seconds()))
}

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}() 


	

	return ch
}