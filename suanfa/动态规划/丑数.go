/**
 * @author X
 * @date 2023/6/21
 * @description
 **/
package main

import "fmt"

//我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i < n+1; i++ {
		num2 := dp[p2] * 2
		num3 := dp[p3] * 3
		num5 := dp[p5] * 5
		dp[i] = min(min(num2, num3), num5)
		if num2 == dp[i] {
			p2++
		}
		if num3 == dp[i] {
			p3++
		}
		if num5 == dp[i] {
			p5++
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
