package main

import "fmt"

func main() {
	nums := []int{100,4,200,1,3,2,6}
	fmt.Println(longestConsecutive(nums))
}


func longestConsecutive(nums []int) int {
    mapa := make(map[int]bool)

	if len(nums) == 0{
		return 0
	} 


	for _, num := range nums {
		mapa[num] = true
	}

	maxl := 0
	for num := range mapa {

		if ! mapa[num - 1] {
			l := 1
			cur := num


			for mapa[cur + 1] {
				cur += 1
				l ++
			}

			maxl = max(maxl, l)
		}	
	}

	return maxl
}

// func longestConsecutive(nums []int) int {
//     mapa := make(map[int]struct{})

// 	if len(nums) == 0{
// 		return 0
// 	} 


// 	for _, num := range nums {
// 		mapa[num] = struct{}{}
// 	}

// 	maxlen := 1

// 	for i := 0; i < len(nums); i++ {
// 		l := 1
// 		num := nums[i]
// 		for {
// 			if _, ok := mapa[num - 1]; ok {
// 				num -= 1
// 				l += 1
// 			}else {
// 				maxlen = max(maxlen, l)
// 				break
// 			}
// 		}
// 	}

// 	return maxlen
	
// }