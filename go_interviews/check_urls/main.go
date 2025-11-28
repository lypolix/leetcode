package main

import (
	"fmt"
	"sync"
)


var links =[]string{}


func main() {
	count := workerpool(10, links)
	fmt.Println(count)
}

func workerpool(workers int, links []string) int {
	count := 0
	mu := &sync.Mutex{}
	linksChan := make(chan string, len(links))
	for _, link := range links {
		linksChan <- link
	}
	close(linksChan)
	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i ++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range linksChan{
				if err := checkUrl(link); err == nil {
					mu.Lock()
					count ++
					mu.Unlock()
				}
			}
		}()
	}
	wg.Wait()
	return count
}


func checkUrl(url string) (err error) {

	return err
}