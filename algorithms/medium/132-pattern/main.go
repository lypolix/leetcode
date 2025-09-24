package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(find132pattern(nums))
}

func find132pattern(nums []int) bool {
    
	stack := []int{}
	middle := -int(^uint(0) >> 1) - 1

	for i := len(nums) - 1; i >= 0; i -- {
		if nums[i] < middle {
			return true
		}

		for len(stack) > 0 && nums[i] > stack[len(stack) - 1] {
			middle = max(middle, stack[len(stack) - 1])
			stack = stack[:len(stack) - 1]
		}

		stack = append(stack, nums[i])

	}
	return false
	
}