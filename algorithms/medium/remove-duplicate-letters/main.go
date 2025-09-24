package main

import (
	"fmt"
	
)


func main() {
	s := "bcabc"
	fmt.Println(removeDuplicateLetters(s))
}
func removeDuplicateLetters(s string) string {
	mapa := make([]int, 26)
	for i := 0; i < len(s); i ++ {
		mapa[int(s[i] - 'a')] ++
	}

	in := make([]bool, 26)
	res := make([]byte, 0, len(s))


	for i := 0; i < len(s); i ++ {
		mapa[int(s[i] - 'a')] --

		if in[int(s[i] - 'a')]{
			continue
		}

		for len(res) > 0 && res[len(res) - 1] > s[i] && mapa[int(res[len(res) - 1] - 'a')] > 0 {
			in[int(res[len(res) - 1] - 'a')] = false
			res = res[:len(res) - 1]
		}

		res = append(res, s[i])
		in[int(s[i] - 'a')] = true 
	}

	return string(res)
}