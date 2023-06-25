package main

import "fmt"

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step := 0
	curDistance := 0  // 当前覆盖最远距离
	nextDistance := 0 // 下一步覆盖最远距离
	for i := 0; i < len(nums); i++ {
		if i+nums[i] > nextDistance {
			nextDistance = nums[i] + i
		}
		if i == curDistance {
			if curDistance < len(nums)-1 {
				step++
				curDistance = nextDistance
				if nextDistance >= len(nums)-1 {
					break
				}
			} else {
				break
			}
		}
	}
	return step
}

func jump2(nums []int) int {
	step := 0
	curDistance := 0  // 当前覆盖最远距离
	nextDistance := 0 // 下一步覆盖最远距离
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > nextDistance {
			nextDistance = nums[i] + i
		}
		if i == curDistance {
			step++
			curDistance = nextDistance
		}
	}
	return step
}

func main() {
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
	fmt.Println(jump(nums))
	fmt.Println(jump2(nums))
}
