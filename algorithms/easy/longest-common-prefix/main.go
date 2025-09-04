package main

import "fmt"


func main() {
	fmt.Printf(longestCommonPrefix([]string{"flower","flow","flight"})) //пример
}

func longestCommonPrefix(strs []string) string {
	j := 0
	if len(strs) == 0{
		return ""
	}
	answ := ""
    for 1 == 1 {
		if (len(strs[0]) <= j){
			return answ
		}
		k := strs[0][j]
		for i := 1; i < len(strs); i ++ {
			if len(strs[i]) <= j{
				return answ
			}else if strs[i][j] != k{
				return answ 
			}
		}
		answ += string(strs[0][j])
		j += 1
	}
	return answ
}
