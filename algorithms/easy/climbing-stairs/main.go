package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
    cache := make([]int, n + 1)
    for i := 0; i <= n; i++ {
        cache[i] = -1
    }
    return Run(0, n, cache)
}

func Run (step int, n int, cache []int) int {
    if step > n{
        return 0
    }
    if step == n {
        return 1
    }

    if cache[step] != -1 {
        return cache[step]
    }
    cache[step] = Run(step + 1, n, cache) + Run(step + 2, n, cache) 
    return cache[step] 
}
