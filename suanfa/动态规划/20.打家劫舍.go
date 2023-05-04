package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return _max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = _max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = _max(dp[i-1], dp[i-2]+nums[i])
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

//func _max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func main() {
//	nums := []int{2, 7, 9, 3, 1}
//	res := rob(nums)
//	fmt.Println(res)
//}
