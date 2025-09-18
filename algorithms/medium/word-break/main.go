package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet","code"}
	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	mapa := make(map[string]bool)
	for _, word := range wordDict {
		mapa[word] = true
	}
	bools := make([]bool, len(s) + 1)
	bools[0] = true
    for i := 1; i <= len(s); i ++ {
		
		for j := i - 1; j >= 0; j -- {
			if bools[j] && mapa[s[j:i]] {
				bools[i] = true
				break
			}  
		}

	}
	return bools[len(s)] 
}