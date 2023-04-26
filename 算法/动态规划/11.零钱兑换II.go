/**
 * @author X
 * @date 2023/3/27
 * @description
 **/
package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func main() {
	coins := []int{2}
	amount := 3
	fmt.Println(change(amount, coins))

}
