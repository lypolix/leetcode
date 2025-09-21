package main

import (
	"fmt"
	"sort"
	"strings"
)


func main () {
	s := "tree"
	fmt.Println(frequencySort(s))
}

func frequencySort(s string) string {
    mapa := make(map[string]int)

	list := []string{}

	for i := 0; i < len(s); i++ {
		mapa[string(s[i])] ++
	}

	for key, _ := range mapa {
		list = append(list, key)
	}

	sort.Slice(list, func(i, j int) bool {
		return mapa[list[i]] > mapa[list[j]]
	})

	res := ""

	for i := 0; i < len(list); i ++ {
		res += strings.Repeat(list[i], mapa[list[i]])
	}

	return res


}