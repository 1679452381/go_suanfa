/**
 * @author X
 * @date 2023/4/2
 * @description
 **/
package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			if dp[j] > dp[j-i*i]+1 {
				dp[j] = dp[j-i*i] + 1
			}
		}
	}
	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}
func main() {
	n := 12
	fmt.Println(numSquares(n))
}
