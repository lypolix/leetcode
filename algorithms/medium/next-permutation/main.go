package main

import "fmt"


func main() {
	nums := []int{1,1,5}
	nextPermutation(nums)
	fmt.Println(nums)
} 


func nextPermutation(nums []int)  {
	mn := 0
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
    for i := len(nums) - 2; i >= 0; i -- {
		if nums[i] < nums[i + 1] {
			mn = i + 1
			for j := i + 1; j < len(nums); j++ {
				if nums[mn] >= nums[j] && nums[j] > nums[i] {
					mn = j
				}
			}
			nums[i], nums[mn] = nums[mn], nums[i]
			for j := i + 1; j < i + 1 + (len(nums) - i)/2; j ++ {
				nums[j], nums[len(nums) - j + i] = nums[len(nums) - j + i], nums[j]
			}
			return 
		}
	}

	for j := 0; j < len(nums)/2; j ++ {
		nums[j], nums[len(nums) - j -1] = nums[len(nums) - j - 1], nums[j]
	}

}
