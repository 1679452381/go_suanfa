/**
 * @author X
 * @date 2023/3/22
 * @description
 **/
package main

import "fmt"

func _max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = _max(dp[i], _max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}

func main() {
	n := 13
	res := integerBreak(n)
	fmt.Println(res)
}
