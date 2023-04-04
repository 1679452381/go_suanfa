/**
 * @author X
 * @date 2023/3/28
 * @description
 **/
package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}
