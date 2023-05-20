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
			dp[i] = _max(dp[i], _max(dp[i-j]*j, (i-j)*j)) //[0 1 1 2 4 6 9 12 18 27 36 54 81 108]
			//[0 1 2 3 4 5 6 7  8  9  10 11 12 13 ]
			//dp[i] = _max(dp[i-j]*j, (i-j)*j)            //[0 0 0 2 4 6 9 12 16 24 30 45 54 72]

		}
	}
	fmt.Println(dp)
	println(dp)
	return dp[n]
}

func main() {
	n := 13
	res := integerBreak(n)
	fmt.Println(res)
}
