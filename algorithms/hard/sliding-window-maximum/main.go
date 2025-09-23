package main

import "fmt"

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

func maxSlidingWindow(nums []int, k int) []int {
    queue := []int{}
	res := make([]int, 0, len(nums) - k + 1)
	if len(nums) == 0 || k == 0 {
		return queue
	}

	for i := 0; i < len(nums); i ++ {
		for len(queue) > 0 && nums[queue[len(queue) - 1]] <= nums[i] {
			queue = queue[:len(queue) - 1]
		}
		queue = append(queue, i)

		if queue[0] <= i - k {
			queue = queue[1:]
		}

		if i - k >= -1 {
			res = append(res, nums[queue[0]])
		}
	}
	
	return res
}