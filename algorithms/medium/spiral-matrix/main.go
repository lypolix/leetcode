package main

import "fmt"

func main() {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(spiralOrder(matrix))
}


func spiralOrder(matrix [][]int) []int {
    
	res := []int{}
	if len(matrix) == 0 {
		return []int{}
	}
	left, right := 0, len(matrix[0]) - 1
	top, botton := 0, len(matrix) - 1


	for left <= right && top <= botton {
		for i := left; i <= right; i ++ {
			res = append(res, matrix[top][i])
		}
		top ++

		for i := top; i <= botton; i ++ {
			res = append(res, matrix[i][right])
		}	
		right --

		if top <= botton {
			for i := right; i >= left; i -- {
				res = append(res, matrix[botton][i])
			}
			botton --
		}

		if left <= right {
			for i := botton; i >= top; i -- {
				res = append(res, matrix[i][left])
			}
			left ++
		}


	}

	return res


}
