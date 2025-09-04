package main

func main() {
	singleNumber([]int{4,1,2,1,2})
}

func singleNumber(nums []int) int {
    dict := make(map[int]int)
    for _, num := range nums{
        dict[num] += 1
    }

    for k, v := range dict{
        if v == 1{
            return k
        }
    }

    return 0
}