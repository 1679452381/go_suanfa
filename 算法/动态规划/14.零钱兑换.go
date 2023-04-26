/**
 * @author X
 * @date 2023/3/29
 * @description
 **/
package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j] > dp[j-coins[i]]+1 {
				dp[j] = dp[j-coins[i]] + 1
			}
		}
	}
	//fmt.Println(dp)
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
func main() {
	coins := []int{1, 4, 9}
	amount := 12
	fmt.Println(coinChange(coins, amount))
}
