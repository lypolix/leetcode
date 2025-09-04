package main

func main() {
	minKBitFlips([]int{0,1,0}, 1)
}


func minKBitFlips(nums []int, k int) int{
	n := len(nums)
	flag := 0
	count := 0
	for i:= 0; i < n; i ++ {
		if i >= k && nums[i - k] == 2{
			flag ^= 1

		}

		if nums[i] == flag {
			if i > n - k {
				return - 1
			} 
			count += 1
			flag ^= 1
			nums[i] = 2
		}
	} 
	return count
}
