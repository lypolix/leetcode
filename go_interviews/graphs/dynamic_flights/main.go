package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)


func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, k, s, f int
	fmt.Fscan(in, &n, &m, &k, &s, &f)
	
	graph := make([][][2]int, n + 1)
	var a, b, c int
	for i := 0; i < m; i ++ {
		fmt.Fscan(in, &a, &b, &c)
		graph[a] = append(graph[a], [2]int{b, c})
	}
				
	const inf = math.MaxInt64

	dp := make([][]int, n + 1)
	for i := range dp {
		dp[i] = make([]int, k + 1)
		for j := 0; j <= k; j ++ {
			dp[i][j] = inf
		}
	}
	dp[s][0] = 0

	for steps := 1; steps <= k; steps++ {
		for node := 1; node <= n; node ++ {
			if dp[node][steps - 1] ==inf {
				continue
			}
			
			for _, fl := range graph[node] {
				next := fl[0]
				cost := fl[1]
				if dp[node][steps - 1] + cost < dp[next][steps] {
					dp[next][steps] = dp[node][steps - 1] + cost
				}
			}
		}
	}

	ans := inf

	for steps := 0; steps <= k; steps++ {
		if dp[f][steps] < ans {
			ans = dp[f][steps]
		}
	}

	if ans == inf {
		fmt.Println(-1)
	}else {
		fmt.Println(ans)
	}
	
}

