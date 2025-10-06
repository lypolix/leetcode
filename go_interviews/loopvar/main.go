package main

import "fmt"

func main() {
	numbers := []*int{}
	for i := 0; i < 5; i ++ {
		i := i
		numbers = append(numbers, &i)
	}

	for _, number := range numbers {
		fmt.Printf("%d", *number)
	}
}
