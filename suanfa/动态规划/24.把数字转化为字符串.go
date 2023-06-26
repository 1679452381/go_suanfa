package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	// 转成字符串
	numStr := strconv.Itoa(num)
	if len(numStr) < 2 {
		return 1
	}
	fmt.Println(numStr)
	n := len(numStr)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		str := numStr[i-2 : i]
		fmt.Println(str)
		if str >= "10" && str <= "25" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	fmt.Println(dp)
	return dp[n]
}
func main() {
	fmt.Println(translateNum(12258))
}
