package main

import (
	"fmt"
	
)


func main() {
	s := "bcabc"
	fmt.Println(removeDuplicateLetters(s))
}
func removeDuplicateLetters(s string) string {
	res := ""

	mapa := make([]int, 26)

	for i := 0; i < 26; i ++ {
		mapa[i] = 0
	}

	for _, val := range s {
		if mapa[int((val) - 'a')] != 0 {
			continue
		} else {
			mapa[int((val) - 'a')] = 1
		}
	}

	for i := 0; i < 26; i ++ {
		if mapa[i] == 1{
			res += string(byte(int('a') + i))
		}
	}

	return res
}
