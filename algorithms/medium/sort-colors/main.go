package main

import "fmt"


func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int)  {
	low := 0 
	med := 0
	high := len(nums) - 1
	for med < high{
		if nums[med] == 0 {
			nums[low], nums[med] = nums[med], nums[low]
			low ++
			med ++
		} else if nums[med] == 1 {
			med ++
		} else {
			nums[high], nums[med] = nums[med], nums[high]
			high --
		}
	}
}
