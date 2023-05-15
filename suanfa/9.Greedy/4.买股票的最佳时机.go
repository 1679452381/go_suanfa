package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	max := 0
	for i := 0; i < len(prices); i++ {
		profit := prices[i] - minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		if profit > max {
			max = profit
		}
	}
	return max
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
