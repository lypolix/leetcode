package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	graph := make([][]int, n + 1)
	for i := 0; i < m; i ++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n + 1)
	concomps := make([][]int, 0)

	queue := make([]int, 0)

	for q := 1; q <= n; q ++ {
		if visited[q] {
			continue
		}
		comp := make([]int, 0)
		queue = queue[:0]
		queue = append(queue, q)

		visited[q] = true

		for head := 0; head < len(queue); head ++ {
			u := queue[head]
			comp = append(comp, u)
			for _, w := range graph[u] {
				if !visited[w] {
					visited[w] = true
					queue = append(queue, w)
				}
			}
		}
		concomps = append(concomps, comp)
	}

	fmt.Fprintln(out, len(concomps))
	for _, comp := range concomps {
		fmt.Fprintln(out, len(comp))
		for i, v := range comp {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
	}
	
}


