package main

import (
	"fmt"
	"sort"
)

func calcRes(goods []int, buyerNeeds[] int) int {
	res := 0
	sort.Slice(goods, func(i, j int) bool {
		return goods[i] < goods[j]
	})

	for i := 0; i < len(buyerNeeds); i ++ {
	    diff := (buyerNeeds[i] - binsearch(goods, buyerNeeds[i]))
		if diff >= 0 {
			res += diff
		}else {
			res -= diff
		}
	}

	return res																														

}

func binsearch(mass []int, need int) int{
	left := 0
	right := len(mass) - 1

	for left < right - 1{
		mid := (right - left)/2 + left
		if mass[mid] < need{
			left = mid
		} else {
			right = mid 
		}

	}
	if left == right {
		return mass[left]
	} else {
		if need - mass[left] < mass[right] - need {
			return mass[left]
		} else {
			return mass[right]
		}
	}
}


func main() {
	goods := []int{1, 2, 5, 8, 9}
	needs := []int{1, 2, 6, 8, 12}															
	fmt.Println(calcRes(goods, needs))
}
