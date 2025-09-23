package main

import "fmt"

func main() {
	s := "abcde"
	words := []string{"a","bb","acd","ace"}

	fmt.Println(numMatchingSubseq(s, words))
}

type wordState struct {
	word string
	index int
}

func numMatchingSubseq(s string, words []string) int {
    count := 0

	buckets := make(map[byte][]wordState)

	for _, w := range words {
		c := w[0]
		buckets[c] = append(buckets[c], wordState{word:w, index: 0})
	}

	for i := 0; i < len(s); i ++ {

		oldBacket := buckets[s[i]]
		buckets[s[i]] = nil

		for _, ws :=  range oldBacket {
			ws.index ++ 
			if ws.index == len(ws.word) {
				count ++
			} else {
				buckets[ws.word[ws.index]] = append(buckets[ws.word[ws.index]], ws)
			}
		}
	}

	return count 
}
