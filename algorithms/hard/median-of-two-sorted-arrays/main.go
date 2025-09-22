package main

import (
	"fmt"
	"math"
)


func main() {
	nums1 := []int{1,3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
		
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low := 0
	high := m


	for low <= high {
		i := (high + low)/2
		j := (n + m + 1)/2 - i

		maxLeft1 := math.Inf(-1)
		if i != 0 {
			maxLeft1 = float64(nums1[i - 1])
		}
		minRight1 := math.Inf(1)
		if i != m {
			minRight1 = float64(nums1[i])
		}

		maxLeft2 := math.Inf(-1)
		if j != 0 {
			maxLeft2 = float64(nums2[j - 1])
		} 

		minRight2 := math.Inf(1)
		if j != n {
			minRight2 = float64(nums2[j]) 
		}

		if minRight1 >= maxLeft2 && minRight2 >= maxLeft1 {
			if (n + m) % 2 == 0 {
				return (math.Max(maxLeft1, maxLeft2) + math.Min(minRight1, minRight2))/2
			} else {
				return math.Max(maxLeft1, maxLeft2)
			}
		} else if minRight1 < maxLeft2 {
			low = i + 1

		} else {
			high = i - 1
		}
	}

	return 0.0
}
