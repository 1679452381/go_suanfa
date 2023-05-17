package main

import (
	"fmt"
)

func candy(ratings []int) int {
	res := make([]int, 0, len(ratings))
	res = append(res, 1)
	fmt.Println(ratings)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res = append(res, res[i-1]+1)
		} else {
			res = append(res, 1)
		}
	}
	fmt.Println(res)
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if res[i] < res[i+1]+1 {
				res[i] = res[i+1] + 1
			}
		}
	}
	fmt.Println(res)
	sum := 0
	for i := 0; i < len(res); i++ {
		sum += res[i]
	}
	return sum
}
func main() {
	//ratings := []int{1, 2, 2, 5, 4, 3, 2}
	ratings := []int{1, 2, 87, 87, 87, 2, 1}
	//[1,2,3,1,1,1,1,1]
	//[1,2,3,1,3,2,1,]
	//ratings := []int{1, 0, 2}
	fmt.Println(candy(ratings))
}
