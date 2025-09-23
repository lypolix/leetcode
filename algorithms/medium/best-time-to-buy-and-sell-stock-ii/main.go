package main

import "fmt"


func main() {
	prices := []int{7,1,5,3,6,4}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
    answ := 0
	if len(prices) <= 1 {
		return 0
	}

	for i := 1; i < len(prices); i ++ {
		if prices[i] > prices[i - 1]{
			answ += prices[i] - prices[i - 1]
		}
	}
	return answ
}