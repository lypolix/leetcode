package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int

	fmt.Fscan(reader, &n)

	g := make([][]int, n)
	for i := 0; i < n; i ++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j ++{
			fmt.Fscan(reader, &g[i][j])
		}
	}

	for i := 0; i < n; i ++ {
		for j := 0; j < n; j ++ {
			for k := 0; k < n; k ++  {
				if (g[i][j] == 1 && g[j][k] == 1 && g[i][k] == 0) {
					fmt.Println("NO")
					return
				}
				
			}
		}
	}

	fmt.Println("YES")
}
