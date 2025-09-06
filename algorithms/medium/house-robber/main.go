package main

import "fmt"

func main () {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	l := len(nums)
    prev := make([]int, l)
	for i := l - 1; i >= 0;i --{
		prev[i] = nums[i]
		if i + 3  <= l - 1{
			prev[i] += max(prev[i + 2], prev[i + 3])
		}else if i + 2 <= l - 1 {
			prev[i] += prev[i + 2]
		}
	}
	if l > 1 {
		return max(prev[0], prev[1])
	}else {
		return prev[0]
	}
	
}
