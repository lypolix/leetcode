package main

import "fmt"

func main()  {
	word1 := "sea"
	word2 := "eat"

	fmt.Println(minDistance(word1,word2))
	
}

func minDistance(word1 string, word2 string) int {
	maxL := 0
	dp := make([][]int, len(word1) + 1)
	for i := 0; i <= len(word1); i++  {
		dp[i] = make([]int, len(word2) + 1)
		dp[i][0] = 0 
	}

	for i := 0; i <= len(word2); i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= len(word1); i ++ {
		for j := 1; j <= len(word2); j ++ {
			if word1[i - 1] == word2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}else {  
				if dp[i][j - 1] > dp[i - 1][j] {
					dp[i][j] = dp[i][j - 1]
				} else {
					dp[i][j] = dp[i - 1][j]
				}
			}
		}
	}

	maxL = dp[len(word1)][len(word2)]


	answ := len(word1) + len(word2) - 2 * maxL
	return answ


}