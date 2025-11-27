package main

import (
	"bufio"
	"fmt"
	"os"
)


func main () {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	g := make([][]bool, n)

	for i := range g {
		g[i] = make([]bool, n)
	}

	for i := 0; i < m; i ++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		a --
		b --
		if a != b {
			g[a][b] = true
			g[b][a] = true
		}
	}

	full := true

	for i := 0; i < n; i ++ {
		for j := i + 1; j < n; j ++ {
			if !g[i][j]  {
				full = false
				break
			}
		} 
		if !full {
			break
		}
	}

	if full {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
