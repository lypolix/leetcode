package main

import (
	"fmt"
	"sort"
)


func main() {

	strs := []string{"eat","tea","tan","ate","nat","bat"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
    hashmap := make(map[string][]string)

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool{
			return runes[i] < runes[j]
		})
		sortedStr := string(runes)

		hashmap[sortedStr] = append(hashmap[sortedStr], str)
	}

	result := make([][]string, 0, len(hashmap))

	for _, val := range hashmap {
		result = append(result, val)
	}

	return result
}
