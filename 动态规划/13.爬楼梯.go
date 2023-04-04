/**
 * @author X
 * @date 2023/3/29
 * @description
 **/
package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 利用01背包问题思想解决爬楼梯问题
func climbStairsBy01(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	m := 2
	for j := 2; j <= n; j++ {
		for i := 1; i <= m; i++ {
			dp[j] += dp[j-i]
		}
	}

	return dp[n]
}
func main() {
	n := 10
	fmt.Println(climbStairs(n))
	fmt.Println(climbStairsBy01(n))
}
