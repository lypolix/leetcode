package main

import (
	"fmt"
	"net/http"
	"sync"
	"context"
)

func worker(ctx context.Context, in <-chan string, out chan<- string) {
	select {
	case <-ctx.Done():
		return
	case url, ok := <-in:
		if !ok {
			return
		}
		resp, err := http.Get(url)
		if err != nil {
			out <- fmt.Sprintf("%s - is not ok", url)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			out  <- fmt.Sprintf("%s - is ok", url)
		}else {
			out <- fmt.Sprintf("%s - is not ok", url)
		}
	}																						
}

type result struct {
	url string
	res *http.Response
	err error
}

func main() {
	var urls = []string {
		"http://ooo.ru",
		"http://ooo.ru",
		"http://ooo.ru",
		"http://ooo.ru",
		"http://ooo.ru",
	}

	urlCh := make(chan string)

	go func() {
		for _, url := range urls {
			urlCh <- url
		}
		close(urlCh)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}
	results := make(chan string)

	for range 100 {
		wg.Add(1)
		go func () {
			defer wg.Done()
			worker(ctx, urlCh, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	
	for r := range results {
		fmt.Println(r)
	}
		
}
