package main

import "fmt"

func main () {
	candidates := []int{2,3,6,7}
	target := 7

	fmt.Println(combinationSum(candidates, target))
}

func combinationSum(candidates []int, target int) [][]int {
    
	res := [][]int{}
	backtracking := func(*[][]int, []int, int, int){}
	list := []int{}
	sum := 0
	backtracking = func(res *[][]int, list []int, sum int, j int) {

		for i := j; i < len(candidates); i++ {
			if sum + candidates[i] == target {
				list = append(list, candidates[i])
				tmp := make([]int, len(list))
				copy(tmp, list)
				*res = append(*res, tmp)
				list = list[:len(list) - 1]
			} else if sum + candidates[i] < target {
				list = append(list, candidates[i])
				backtracking(res, list, sum + candidates[i], i)
				list = list[:len(list) - 1]
			}
		}
	}

	backtracking(&res, list, sum, 0)

	return res
}