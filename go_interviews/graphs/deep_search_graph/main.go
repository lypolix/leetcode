package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)


func dfs(v int, g [][]int, used []bool, comp *[]int){
	used[v] = true
	*comp = append(*comp, v + 1)
	for _, to := range g[v] {
		if !used[to] {
			dfs(to, g, used, comp)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	g := make([][]int, n)
	for i := 0; i < m; i ++ {																				
		var a, b int
		fmt.Fscan(reader, &a, &b)
		a --
		b --
		g[a] = append(g[a], b)
		g[b] = append(g[b], a	)
	}


	used := make([]bool, n)
	var comp []int
	dfs(0, g, used, &comp)

	sort.Ints(comp)
	fmt.Println(len(comp))
	for i, v := range comp{
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()


}
