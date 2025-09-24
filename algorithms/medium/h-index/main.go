package main

import "fmt"

func main() {
	citations := []int{3,0,6,1,5}
	fmt.Println(hIndex(citations))
}

func hIndex(citations []int) int {
    col := make([]int, len(citations) + 1)

	for i := 0; i < len(citations); i ++ {
		if citations[i] > len(citations){
			col[len(citations)] ++
		} else {
			col[citations[i]] ++ 
		}
	}

	total := 0  
	for i := len(citations); i >= 0; i -- {
		total += col[i]
		if total >= i {
			return i
		}
	}

	return 0
}
