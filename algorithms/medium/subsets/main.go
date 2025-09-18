package main

import "fmt"



func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

func subsets(nums []int) [][]int {
    res := [][]int{}
	var curr []int

	

	var backtrack func(start int)
	backtrack = func(start int) {

		tmp := make([]int, len(curr))
		copy(tmp, curr)
		res = append(res, tmp)


		for i := start; i < len(nums); i ++ {
			curr = append(curr, nums[i])
			backtrack(i + 1)
			curr = curr[:len(curr)-1]
		}

	}
	backtrack(0)

	return res

}