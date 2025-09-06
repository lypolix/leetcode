package main

import "fmt"

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
    steps := [2]int{m - 1, n - 1}
	cash := make(map[[2]int]int)
	return rec(steps, 0, cash)
    
}

func rec(steps [2]int, count int, cash map[[2]int]int) int{
	if result, ok := cash[steps]; ok {
		return result
	}
    if steps[0] + steps[1] == 0{
        return 1
    }

	if steps[0] != 0 && steps[1] != 0 {
        steps1 := [2]int{steps[0] - 1, steps[1]}
		steps2 := [2]int{steps[0], steps[1] - 1}
		count ++
		result := rec(steps1, count, cash) + rec(steps2, count, cash)
		cash[steps] = result
        return result
    }else if steps[0] != 0 {
        steps1 := [2]int{steps[0] - 1, steps[1]}
		count ++
		result := rec(steps1, count, cash)
		cash[steps] = result
        return result
    }else if steps[1] != 0 {
        steps1 := [2]int{steps[0], steps[1] -1}
		count++
		result := rec(steps1, count, cash)
		cash[steps] = result
        return result
    } 
	return 0
}