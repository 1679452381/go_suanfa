/**
 * @author X
 * @date 2023/4/4
 * @description
 **/
package main

import "fmt"

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return _max(nums[0], nums[1])
	}
	nums1 := nums[0 : len(nums)-1]
	nums2 := nums[1:len(nums)]
	res1 := robRange(nums1)
	res2 := robRange(nums2)
	if res1 > res2 {
		return res1
	}
	return res2
}
func robRange(nums []int) int {
	dp := make([]int, len(nums))
	fmt.Println(dp, nums)
	dp[0] = nums[0]
	dp[1] = _max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = _max(dp[i-1], dp[i-2]+nums[i])
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}

func main() {
	nums := []int{1, 2, 3, 1, 3, 4}
	fmt.Println(rob2(nums))
}
