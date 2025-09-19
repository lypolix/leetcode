package main

import "fmt"


func main() {
	arr:= []int{3,2,4,1}
	fmt.Println(pancakeSort(arr))
}

func flip (arr []int, k int) {
	for i, j := 0, k - 1; j > i; i, j = i + 1, j - 1 {
		arr[i], arr[j] = arr[j], arr[i]
	}	
}

func pancakeSort(arr []int) []int {
    
	res := []int{}

	for size := len(arr); size > 1; size-- {
		
		maxIdx := 0
		for i := 1; i < size; i ++ {
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}

		if maxIdx == size - 1 {
			continue
		}

		if maxIdx != 0 {
			flip(arr, maxIdx + 1)
			res = append(res, maxIdx + 1)
		}

		flip(arr, size)
		res = append(res, size)
	}

	return res
}