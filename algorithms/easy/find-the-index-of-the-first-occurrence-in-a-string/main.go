package main

import "fmt"

func main() {
	
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	count := len(needle)
	k := 0
    for i := 0; i < len(haystack); i ++ {
		if haystack[i] == needle[k]{
			k ++
			if k == count{
				return i - count + 1
			}
		}else if k > 0{
			i -= k
			k = 0
		}
	}
	return -1
}