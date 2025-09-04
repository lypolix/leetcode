package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	result := permute(nums)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
    
	if len(nums) == 0 {
		return [][]int{}
	} 

	var result [][]int
	var generate func(int)

	generate = func(index int) {
		if index == len(nums) - 1{
			temp := make([]int, len(nums))
			copy(temp, nums)

			result = append(result, temp)
			return
		}

		for i := index; i < len(nums); i ++ {
			nums[i], nums[index] = nums[index], nums[i]
			generate(index + 1)

			nums[index], nums[i] = nums[i], nums[index]
		}
	}

	generate(0)
	return result

}