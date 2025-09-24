package main

import "fmt"


func main() {
	s1 := "great"
	s2 := "rgeat"
	fmt.Println(isScramble(s1, s2))
}

func isScramble(s1 string, s2 string) bool {
    memo := make(map[string]bool)
	return helper(s1, s2, memo)
}

func helper(s1 string, s2 string, memo map[string]bool) bool {
	key := s1 + "#" + s2
	if v, ok := memo[key]; ok {
		return v
	}

	if s1 == s2 {
		memo[key] = true
		return true
	}

	if !isSame(s1, s2) {
		memo[key] = false
		return false
	}
	n := len(s1)	
	for i := 1; i < len(s1); i ++ {
		if helper(s1[:i], s2[:i], memo) && helper(s1[i:], s2[i:], memo) {
			memo[key] = true
			return  true
		}
		

		if helper(s1[:i], s2[n - i:], memo) && helper(s1[i:], s2[:n - i], memo) {
			memo[key] = true
			return true
		}
	}
	memo[key] = false
	return false
}



func isSame(s1 string, s2 string) bool {
	sp := make([]int, 26)
	for i := 0; i < len(s1); i ++ {
		sp[int(s1[i] -'a')] ++
		sp[int(s2[i] -'a')] --
	}

	for i:= 0; i < 26; i ++ {
		if sp[i] != 0 {
			return false
		}
	}

	return true
}
