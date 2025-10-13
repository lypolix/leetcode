package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func getDiscount() (float64){
	resp, err := http.Get("http://discount.com/my")
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	discountStr := strings.TrimSpace(string(body))
	discount, err := strconv.ParseFloat(discountStr, 64)
	if err != nil {
		return 0
	}
	return discount
}


func main () {
	ctx := context.Background()
	v, err := getDiscountWithContext(ctx)
	if err != nil {
		fmt.Println("Что-то пошло не так", err.Error())
		return
	}
	fmt.Printf("Ваша скидка: %v", v)
}

func getDiscountWithContext(ctx context.Context) (float64, error){
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()

	ch := make(chan float64)
	go func () {
		ch <- getDiscount()
		defer close(ch)
	}()



	select {
	case <- ctx.Done():
		return 0, ctx.Err()
	case res := <- ch:
		return res, nil
	}
	
}