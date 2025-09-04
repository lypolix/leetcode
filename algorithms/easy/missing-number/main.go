package main

func main() {
	missingNumber([]int{3,0,1})
}

func missingNumber(nums []int) int {
    list := make([]int, len(nums) + 1)
    for _, num := range nums{
        list[num] = 1
    }

    for i := range list{
        if list[i] != 1{
            return i
        }
    }  
    return -1 
}