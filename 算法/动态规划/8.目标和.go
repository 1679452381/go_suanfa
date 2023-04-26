/**
 * @author X
 * @date 2023/3/23
 * @description
 **/
package main

import (
	"fmt"
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if int(math.Abs(float64(sum))) < target {
		return 0
	}
	c := sum + target
	if c%2 == 1 {
		return 0
	}
	bagSize := c / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}
