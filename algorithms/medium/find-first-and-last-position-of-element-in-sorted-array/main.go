package main

import (
	"fmt"
)

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 8

	fmt.Println(searchRange(nums, target))
}


func first(nums []int, target int) int {

	left := 0
	right := len(nums) -1 
	answ := - 1

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			answ = mid
			right = mid - 1 
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return answ
}

func second(nums []int, target int) int {

	left := 0
	right := len(nums) - 1 
	answ := - 1

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			answ = mid
			left = mid + 1 
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return answ
}


func searchRange(nums []int, target int) []int {
	fst := first(nums, target)
	scd := second(nums, target)

	return []int{fst, scd}
}