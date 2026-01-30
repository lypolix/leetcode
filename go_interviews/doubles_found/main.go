package main

import "fmt"

func main() {
	in := []string{"a", "bb", "bb", "aa", "a", "a", "a"} 
	m := make(map[string]int)

	for _, s := range in {
		m[s]++
	}
	out := []string{}
	for s, p := range m {
		if p > 1 {
			out = append(out, s)
		}
	}

	fmt.Println(out)

}