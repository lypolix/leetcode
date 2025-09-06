package main

import "fmt"

func main () {
	digits := []int{}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {

	i := 1

	for digits[len(digits) - i] == 9 {
		digits[len(digits) - i] = 0
		if len(digits) == i {
			digits = append([]int{1}, digits...)
			return digits
		}
		i++
	}

	digits[len(digits) - i] ++
	return digits
}