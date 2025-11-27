package main

import (
	"fmt"
	"reflect"
)

type Statistic struct {
	UserID int
	Steps  int
}

type Reult struct {
	UserIDs []int
	Steps int
}

func getChampion(statistics [][]Statistic) Reult {
	sumSteps := make(map[int]int)
	sumcount := make(map[int]int)
	for _, stat := range statistics {
		for _, userstat := range stat {
			sumSteps[userstat.UserID] += userstat.Steps
			sumcount[userstat.UserID] += 1
		}
	}
	n := len(statistics)
	ans := 0
	vinner := make([]int, 0)
	for key, val := range sumcount{
		if val == n && sumSteps[key] >= ans{
			ans = sumSteps[key] 
			if sumSteps[key] > ans {
				vinner = []int{key}
			}else {
				vinner = append(vinner, key)
			}
		}
	}
	var res Reult

	res.UserIDs = vinner
	res.Steps = ans

	return  res
}

func main () {
	fmt.Println("example 1", reflect.DeepEqual(
		getChampion(
			[][]Statistic{
				{{UserID: 1, Steps: 2000}, {UserID: 2, Steps: 1500}},
				{{UserID: 2, Steps: 1000}},  
			},
		),
		Reult{
			UserIDs: []int{2},
			Steps: 2500,
		},
	))

	
}