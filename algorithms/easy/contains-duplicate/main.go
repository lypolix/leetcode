package main 

func main() {
	containsDuplicate([]int{1,2,3,1})
}

func containsDuplicate (nums []int) bool {
    dict := make(map[int]int)
    for _, num := range nums {
        dict[num] += 1
        if dict[num] > 1{
            return true
        }
    }

    return false
}
