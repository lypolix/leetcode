package main

import "fmt"

func main() {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	qraph := make([][]int, numCourses)
	cursesBefore := make([]int, numCourses)
	count := 0

	for i := 0; i < len(prerequisites); i ++ {
		c1 := prerequisites[i][0]
		c2 := prerequisites[i][1]
		qraph[c1] = append(qraph[c1], c2)
		cursesBefore[c2] ++
	}

	queue := []int{}
	for i := 0; i < numCourses; i ++ {
		if cursesBefore[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		count ++ 
		course := queue[0]
		queue = queue[1:]

		for _, k := range(qraph[course]) {
			cursesBefore[k] --
			if cursesBefore[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	if count == numCourses {
		return true
	} else {
		return false
	}
}
