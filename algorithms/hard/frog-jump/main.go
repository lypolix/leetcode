package main

import "fmt"

func main() {
	stones := []int{0, 1, 2, 3, 4, 8, 9, 11}
	fmt.Println(canCross(stones))
}

//Динамическое решение 
func canCross(stones []int) bool {
	path := make(map[int]map[int]struct{})

	for _, stone := range stones{
		path[stone] = make(map[int]struct{})
	}

	path[0][1] = struct{}{}
	

	for _,  stone := range stones {

		for jump := range path[stone]{

			for step := jump - 1; step <= jump + 1; step ++ {
				if step <= 0 {
					continue
				}
				next := step + stone

				if _, ok := path[next]; ok {
					path[next][step] = struct{}{}
				}
			}
		}
	}
	return len(path[stones[len(stones) - 1]]) > 0


}

//Рекурсивное решение 

// func canCross(stones []int) bool {
// 	cash := make(map[[2]int]int)
// 	res := rec(stones, 1, 1, cash)
// 	if stones[1] != stones[0]+1 {
// 		return false
// 	}
// 	if res > 0 {
// 		return true
// 	}
// 	return false
// }

// func rec(stones []int, ind int, step int, cash map[[2]int]int) int {
// 	if step <= 0 || ind == -1 {
// 		return 0
// 	}
// 	if ind == len(stones)-1 {
// 		return 1
// 	}
// 	if ind >= len(stones) {
// 		return 0
// 	}

// 	a, b, c := stones[ind]+step+1, stones[ind]+step, stones[ind]+step-1
// 	nextstone := []int{a, b, c}
// 	k := ind

// 	jumped := [3]int{-1, -1, -1}
// 	times := 0
// 	for _, stone := range nextstone {
// 		k = ind + 1
// 		if k >= len(stones) {
// 			break
// 		}
// 		for stone >= stones[k] {
// 			if stone == stones[k] {
// 				jumped[times] = k
// 				break
// 			}
// 			k++
// 			if k >= len(stones) {
// 				break
// 			}
// 		}
// 		times++
// 	}
// 	res := 0
// 	if jumped[0] == -1 {
// 	} else if res1, ok := cash[[2]int{jumped[0], step + 1}]; ok {
// 		res += res1
// 	} else {
// 		res1 := rec(stones, jumped[0], step+1, cash)
// 		cash[[2]int{jumped[0], step + 1}] = res1
// 		res += res1
// 	}

// 	if jumped[1] == -1 {
// 	} else if res2, ok := cash[[2]int{jumped[1], step}]; ok {
// 		res += res2
// 	} else {
// 		res2 := rec(stones, jumped[1], step, cash)
// 		cash[[2]int{jumped[1], step}] = res2
// 		res += res2
// 	}
// 	if jumped[2] == -1 {
// 	} else if res3, ok := cash[[2]int{jumped[2], step - 1}]; ok {
// 		res += res3
// 	} else {
// 		res3 := rec(stones, jumped[2], step-1, cash)
// 		cash[[2]int{jumped[2], step - 1}] = res3
// 		res += res3
// 	}

// 	return res

// }
