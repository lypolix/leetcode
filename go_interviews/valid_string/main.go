package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("{{[]}}"))
}



func isValid(s string) bool {
	stack := []rune{}
	
	pairs := map[rune]rune{')':'(', '}':'{', ']':'['}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				return false
			} 

			last := stack[len(s) - 1]
			if last != pairs[char] {
				return false
			}

			stack = stack[:len(stack) - 1]
		}
	}


	return len(stack) == 0
}

































