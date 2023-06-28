/**
 * @author X
 * @date 2023/6/26
 * @description
 **/
package main

import "fmt"

func jump(nums []int) int {
	currDistance := 0
	nextDistance := 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > nextDistance {
			nextDistance = i + nums[i]
		}
		if i == currDistance {
			step++
			currDistance = nextDistance
		}
	}
	return step
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
