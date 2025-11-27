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
	dists := make([][]int, n)
	used := make([]bool, n)
	dist := make([]int, n)
	for i := 0; i < n; i ++ {
		dists[i] = make([]int, 0)
		dist[i] = inf
	}
	dists[s] = append(dists[s], s)
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
					tmp := make([]int, len(dists[v]))
					copy(tmp, dists[v])
					tmp = append(tmp, u)
					dists[u] = tmp
				}
			}
		}

	}
	for i := 0; i < len(dists[f]); i ++ {
		dists[f][i] ++
	}

	if dist[f] == inf {
		fmt.Println(-1)
	} else {
		for i := 0; i < len(dists[f]); i ++ {
			if i != 0 {
				fmt.Print(" ")
			}
			fmt.Print(dists[f][i])
		}
	}
	fmt.Println()
}