package main

import "fmt"

// 思路：计算所有天数的正利润
func maxProfit2(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			max += profit
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(prices))
}
