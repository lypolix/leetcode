package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
    
	chars := make(map[byte]int)

	maxC := 0
	count := 0
	for i := 0; i < len(s); i ++ {
		if _, ok := chars[s[i]]; !ok {
			count ++
			chars[s[i]] = i
		}else {
			count = min(count + 1, i - chars[s[i]])
			chars[s[i]] = i
		}
		maxC = max(count, maxC)
	}
	return maxC
}

