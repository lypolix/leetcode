package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)



func main() {
	in := bufio.NewReader(os.Stdin)
	var n, s, f int
	fmt.Fscan(in, &n, &s, &f)
	s--
	f--

	graph := make([][]int, n)
	for i := 0; i < n; i ++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j ++ {
			var w int
			fmt.Fscan(in, &w)
			graph[i][j] = w
		}
	}
	const inf = math.MaxInt
	dist := make([]int, n)
	used := make([]bool, n)
	for i := 0; i < n; i ++ {
		dist[i] = inf
	}
	dist[s] = 0

	for i := 0; i < n; i ++ {
		v := -1
		for j := 0; j < n; j ++ {
			if !used[j] && (v == -1 || dist[j] < dist[v]) {
				v = j
			}
		}
		if v == -1 || dist[v] == inf {
			break
		}

		used[v] = true

		for u := 0; u < n; u ++ {
			w := graph[v][u]
			if w >= 0 {
				if dist[u] > dist[v] + w {
					dist[u] = dist[v] + w
				}
			}
		}

	}

	if dist[f] == inf {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[f])
	}
}