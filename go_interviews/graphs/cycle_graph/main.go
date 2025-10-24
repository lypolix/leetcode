package main

import (
	"bufio"
	"fmt"
	"os"
)



var (
	g [][]int
	used []int
	parent []int
	n int
	cycleStart, cycleEnd int
)

func dfs(v, p int) bool {
	used[v] = 1

	for to := 0; to < n; to ++ {
		if g[v][to] == 1 {
			if used[to] == 0 {
				parent[to] = v
				if dfs(to, v) {
					return true
				} 
			} else if to != p {
				cycleStart = to
				cycleEnd = v
				return true
			}
		}
	}

	return false


}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	g = make([][]int, n)
	for i:= 0; i < n; i ++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j ++ {
			fmt.Fscan(reader, &g[i][j])
		}
	}


	used := make([]int, n)
	parent = make([]int, n)
	for i := 0; i < n; i ++ {
		used[i] = 0	
		parent[i] = -1
	}
	cycleStart = -1
	cycleEnd = -1

	for i := 0; i < n; i ++ {
		if used[i] == 0 {
			if dfs(i, -1) {
				break
			}
		}
	}


	if cycleStart == -1 {
		fmt.Println("NO")
	} else {
		cycle := []int{}
		v := cycleEnd
		cycle = append(cycle, v + 1)
		for v != cycleStart {
			v = parent[v]
			cycle = append(cycle, v + 1)
		}
		for i, j := 0, len(cycle) - 1; i < j; i, j = i + 1, j - 1 {
			cycle[i], cycle[j] = cycle[j], cycle[i]
		}

		fmt.Println("YES")
		fmt.Println(len(cycle))
		for i, num := range cycle {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(num)
		}

		fmt.Println()
	}
}
