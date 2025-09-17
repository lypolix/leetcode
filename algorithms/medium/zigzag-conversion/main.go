package main

import "fmt"


func main() {
	s := convert("PAYPALISHIRING", 4)
	fmt.Println(s)
}


func convert(s string, numRows int) string {
	s2 := ""
	str := make([]string, numRows)

	if numRows == 1 {
		return s
	}
	f := 0

	vniz := true
	for _, sim := range s {
		str[f] += string(sim)
		if vniz{
			if f == numRows - 1 {
				vniz = false
				f = numRows - 2
			}else{
				f ++
			}
		} else {
			if f == 0 {
				f = 1
				vniz = true
			} else {
				f--
			}
		} 
	}
	for _, st := range str {
		s2 += st
	}
	return s2
}