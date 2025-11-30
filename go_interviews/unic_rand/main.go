package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(uniqRandn(10))
}


func uniqRandn(n int) []int {
	res := make([]int, n)
	uniq := make(map[int]interface{})
	i := 0
	for i < n {
		newVal := rand.Int()
		if _, ok := uniq[newVal]; ok {
			continue
		}
		uniq[newVal] = struct{}{}
		res[i] = newVal
		i ++
		
	}
	return res
}