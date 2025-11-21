package main

import (
	"bufio"
	"fmt"
	"os"
)



func main () {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)

	graph := make([][]int, n)
	for i := 0; i < m; i ++ {
		var a, b int	
		fmt.Fscan(in, &a, &b)
		a--
		b--
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	colors := make([]int, n)
	for i := 0; i < n; i ++{
		colors[i] = -1 
	}

	for i := 0; i < n; i ++ {
		if colors[i] == -1{
			queue := []int{i}
			colors[i] = 0

			for len(queue) > 0 {
				v := queue[0] 
				queue = queue[1:]

				for _, u := range graph[v]{
					if colors[u] == -1 {
						colors[u] = 1 - colors[v]
						queue = append(queue, u)
					} else if colors[u] == colors[v] {
						fmt.Println("NO")
						return
					}
				}
			}
		}
	}

	
	fmt.Println("YES")
	for i := 0; i < n; i ++ {
		if colors[i] == 0 {
			fmt.Print(i + 1, " ")
		}
	}
	fmt.Println()
	
}