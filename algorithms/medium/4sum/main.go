package main

import (
	"fmt"
	"sort"
)



func main() {
	nums := []int{1,0,-1,0,-2,2}
	target := 0
	fmt.Println(fourSum(nums, target))

}

func fourSum(nums []int, target int) [][]int {
    res := [][]int{}

	sort.Slice(nums, func(i, j int) bool{
		return nums[i] < nums[j]
	})
	left := 0
	right := len(nums) - 1
	for i := 0; i < len(nums) - 1; i ++ {
		if i > 0 && nums[i] == nums[i - 1]{
			continue
		}
		for j := i + 1; j < len(nums); j ++ {
			if j > i + 1 && nums[j] == nums[j - 1] {
				continue
			}
			left = j + 1
			right = len(nums) - 1
			sum := nums[i] + nums[j]
			for left < right {
				if sum + nums[left] + nums[right] == target {
					res = append(res, []int{nums[right], nums[left], nums[i], nums[j]})
					left ++
					right --

					for left < right && nums[left] == nums[left - 1] {
						left ++
					} 

					for left < right && nums[right] == nums[right + 1] {
						right --
					} 
				} else if sum + nums[left] + nums[right] < target {
					left ++
				} else {
					right --
				}
			}
		}
	}

	return res
}