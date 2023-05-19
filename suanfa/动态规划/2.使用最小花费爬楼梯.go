package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, 0)
	dp = append(dp, 0, 0)
	minCost := 0
	for i := 2; i <= len(cost); i++ {
		if dp[i-1]+cost[i-1] > dp[i-2]+cost[i-2] {
			minCost = dp[i-2] + cost[i-2]
		} else {
			minCost = dp[i-1] + cost[i-1]
		}
		dp = append(dp, minCost)
	}
	fmt.Println(dp)
	return dp[len(cost)]
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}

	println(minCostClimbingStairs(cost))
}
