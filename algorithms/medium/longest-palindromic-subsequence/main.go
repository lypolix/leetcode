package main

import "fmt"

func main () {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))

   	for i := 0; i < len(s); i ++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
    }

	for length := 2; length <= len(s); length ++ {
		for i := 0; i <= len(s) - length; i ++ {
			j := i + length - 1 
			if s[i] == s[j] {
				if length == 2 {
					dp[i][j] = 2
				}else {
					dp[i][j] = dp[i + 1][j - 1] + 2
				}
			} else {
				dp[i][j] = max(dp[i + 1][j], dp[i][j - 1])
			}
		}
	}
	return dp[0][len(s) - 1]
}