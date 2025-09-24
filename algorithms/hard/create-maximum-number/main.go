package main

import (
	"fmt"
)


func main() {
	nums1 := []int{3,4,6,5}
	nums2 := []int{9,1,2,5,8,3}
	k := 5
	fmt.Println(maxNumber(nums1, nums2, k))
}


func maxNumber(nums1 []int, nums2 []int, k int) []int {
    
	maxSeq := func(nums []int, k int) []int{
		stack := make([]int, 0, k)

		drop := len(nums) - k 

		for _, num := range nums {
			for drop > 0 && len(stack) > 0 && num > stack[len(stack) -  1] {
				stack = stack[:len(stack) - 1] //уменьшаем только длинну или ещё и капасити?
				drop -- 
			}

			stack = append(stack, num)
		}
		return stack[:k]
	}
	greater := func(nums1 []int, i int, nums2 []int, j int) bool{
		for i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
			i ++
			j ++
		}

		if j == len(nums2) {
			return true
		}

		if i == len(nums1) {
			return false
		}

		return nums1[i] > nums2[j]
	}

	merge := func(nums1, nums2 []int) []int {
		result := make([]int, 0, len(nums1) + +len(nums2))
		i, j := 0, 0 

		for i < len(nums1) || j < len(nums2) {
			if greater(nums1, i, nums2, j) {
				result = append(result, nums1[i])
				i ++
			} else {
				result = append(result, nums2[j])
				j ++
			}
		}
		return result
	}

	

	res := make([]int, 0)
	start := max(0, k - len(nums2))
	end := min(k, len(nums1))

	for i := start; i <= end; i ++ {
		sub1 := maxSeq(nums1, i)
		sub2 := maxSeq(nums2, k - i) 
		cand := merge(sub1, sub2)
		if greater(cand, 0, res, 0) {
			res = cand
		}
	}
	return res

}