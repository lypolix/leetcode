package main

import (
	"fmt"
	"sync"
)


func square(a int) int {
	return a*a
}
func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- f(job)
	}
}

func main() {
	workerNums := 4

	jobs := make(chan int, 10)
	result := make(chan int, 10)
	for i := range 10 {
		jobs <- i
	}
	close(jobs)

	wg := sync.WaitGroup{}

	wg.Add(workerNums)
	for i := range workerNums {
		go func() {
			wg.Done()
			worker(i, square, jobs, result)
		}()
	}

	go func () {
		wg.Wait()
		close(result)
	}()
	
	for value := range result {
		fmt.Println(value)
	}
	
	
}
