package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(nums))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(a int, b int) bool {
		if people[a][0] != people[b][0] {
			return people[a][0] > people[b][0]
		} else {
			return people[a][1] < people[b][1]
		}
	})

	res := make([][]int, len(people))

	for _, pep := range people {
		insert(&res, pep[1], pep)
	}
	return res
}

func insert(res *[][]int, Npep int, pep []int) {

	copy((*res)[Npep+1:], (*res)[Npep:])
	(*res)[Npep] = pep
}
