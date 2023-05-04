package main

import "fmt"

// 可以转换为01背包问题

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := make([]int, target+1)
	for i := nums[0]; i < target; i++ {
		dp[i] = nums[0]
	}
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			if dp[j] < dp[j-nums[i]]+nums[i] {
				dp[j] = dp[j-nums[i]] + nums[i]
			}
		}
	}
	return dp[target] == target
}

func main() {
	nums := []int{1, 2, 5}
	res := canPartition(nums)
	fmt.Println(res)
}
