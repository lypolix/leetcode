package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// )

// func main() {

// 	ctx := context.Background()

// 	var urls = []string {
// 		"http://ernj/fejnf",
// 		"http://ernj/fejnf",
// 		"http://ernj/fejnf",
// 		"http://ernj/fejnf",
// 		"http://ernj/fejnf",
// 	}


// 	type result struct {
// 		url string
// 		err error
// 	}


// 	channelSize := 0
// 	urlInput := concurrency.Generator(ctx, urls, channelSize)
// 	workersCount := 0
// 	resCh := concurrency.Start(ctx, workersCount, urlInput, func(currentUrl string) result{
// 		response, err := http.Get(currentUrl)

// 		if err != nil {
// 			return result{
// 				url : currentUrl, 
// 				err : fmt.Errorf("failed %s, error - %v", currentUrl, err), 
// 			}
// 		}

// 		if response.StatusCode != http.StatusOK {
// 			return result{
// 				url: currentUrl, 
// 				err : fmt.Errorf("failed %s with http code %d", currentUrl, response.StatusCode),
// 			}
// 		}

// 		return result{
// 			url: currentUrl,
// 		}
// 	})

// 	for r := range resCh {
// 		fmt.Printf("URL: %s, Error: %v\n", r.url, r.err)
// 	}
// }
