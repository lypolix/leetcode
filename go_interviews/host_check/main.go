package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)



func main() {
	dur := time.Second * 10
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(dur)
	wg := &sync.WaitGroup{}
	
	defer cancel()
	for {			
		select {
		case <- ticker.C:
			mu.RLock()
			for url := range hosts {
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					err := healthCheck(ctx, url)
					if err != nil {
						cancel()
						fmt.Println(err)
						return
					}
				}(url)
			}
			mu.RUnlock()
			go func() {
				wg.Wait()
			}()
		}
	}
}


var hosts = make(map[string]bool)
var mu = &sync.RWMutex{}

func healthCheck(ctx context.Context, url string) error {
	
	resp, err := http.Get(url)
	if err != nil {
		hosts[url] = false
		return err
	}
	mu.Lock()
	resp.Body.Close()
	if resp.StatusCode == 200 {
		hosts[url] = true
	} else {
		hosts[url] = false
	}
	mu.Unlock()

	return nil
	
}