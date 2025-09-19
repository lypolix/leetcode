package main

import "fmt"


func main () {
	arr := []int{4,3,2,1,0}
	fmt.Println(maxChunksToSorted(arr))
}

func maxChunksToSorted(arr []int) int {

	maxEl := 0
	answ := 0

    for i := 0; i < len(arr); i ++ {
		maxEl = max(maxEl, arr[i])
		if maxEl == i {
			answ += 1
		}
	}

	return answ
}
