package main

import "fmt"


func main() {
	s := "AAAAABBBBCBB"
	k := 4
	fmt.Println(characterReplacement(s, k))
}

type Comb struct {
	a int
	b int
}

func characterReplacement(s string, k int) int {
    
	list := make([]int, 26)

	left, right := 0, 0 

	Comb := &Comb{0, 0}
	answ := 0
	for i := right; i < len(s); i ++ {

	
		list[int(s[i] - 'A')]++
		if list[(s[i] - 'A')] >= Comb.a{
			Comb.a = list[int(s[i] - 'A')]
			Comb.b = int(s[i] - 'A')
		}

		for (i - left + 1 - Comb.a) > k {

			if list[int(s[left] - 'A')] != 0 {
				list[int(s[left] - 'A')] --
			}

			if Comb.b == int(s[left] - 'A') {
				Comb.a = list[Comb.b] 
			}
			left ++
		}
		answ = max(answ, i - left + 1)
		fmt.Println(list)
		fmt.Println(Comb)
		fmt.Println(answ)

	}
	return answ
	
}