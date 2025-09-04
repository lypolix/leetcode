package main

func main() {
	findDisappearedNumbers([]int{4,3,2,7,8,2,3,1})
}

func findDisappearedNumbers(nums []int) []int {

    dict := make([]int, len(nums) + 1)
    for _, num := range nums{
        dict[num] = 1
    }
    ans := make([]int, 0)
    for k, v := range dict {
        if v == 0 && k != 0{
            ans = append(ans, k)
        }
    }
    return ans
}
