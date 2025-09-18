package main

import (
	"fmt"
	"strings"
)


func main() {
	nums0 := []string{"A","C","A","B","D","B"}
	str := strings.Join(nums0, "")
	nums := []byte(str)
	fmt.Println(leastInterval(nums, 1))

}
func leastInterval(tasks []byte, n int) int {
    maxK := 0
	hash := make(map[byte]int)
	hash1 := make(map[int]int)
	maxL := 0
	for i := 0; i < len(tasks); i++ {
		hash[tasks[i]] += 1
		maxL = max(maxL, hash[tasks[i]])
		hash1[hash[tasks[i]]] +=1
	}
	maxK = hash1[maxL]

	answ := max(len(tasks), (maxL - 1) * (n + 1) + maxK)
	return answ
}
