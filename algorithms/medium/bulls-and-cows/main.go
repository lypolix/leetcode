package main

import (
	"fmt"
	"strconv"
)

func main() {
	secret := "1807"
	guess := "7810"

	fmt.Println(getHint(secret, guess))
}

func getHint(secret string, guess string) string {
    
	list := make([][]int, 10)
	for i := 0; i < 10; i ++ {
		list[i] = make([]int, 2)
	}

	bulls := 0
	cows := 0
	for i := 0; i < len(secret); i ++ {
		if secret[i] == guess[i] {
			bulls += 1
		} else {
			list[secret[i] - '0'][0] +=1
			list[guess[i] - '0'][1] += 1
		}
	}

	for i := 0; i < 10; i ++ {
		cows += min(list[i][0], list[i][1])
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
