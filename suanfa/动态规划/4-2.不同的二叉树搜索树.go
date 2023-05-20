package main

import "fmt"

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	fmt.Println(dp)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
