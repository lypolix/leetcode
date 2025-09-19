package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 4

	fmt.Println(countAndSay(n))
}

func countAndSay(n int) string {

	answ := "1"

	for i := 1; i < n; i ++ {
		s := answ[0]
		kol := 1
		answ2 := ""
		if len(answ) == 1 {
			answ = "1" + answ
			continue
		}
		for j := 1; j < len(answ); j ++ {
			if answ[j] == s {
				kol += 1
			} else {
				answ2 += strconv.Itoa(kol) + string(answ[j - 1])
				kol = 1
			}
			s = answ[j]
		}

		answ2 += strconv.Itoa(kol) + string(answ[len(answ) - 1])
		answ = answ2
	}
	return answ
}