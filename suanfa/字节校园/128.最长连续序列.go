package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	max := 0
	for key := range numSet {
		//跳过外层重复循环，降低时间复杂度
		if !numSet[key-1] {
			curr := 1
			for numSet[key+1] {
				curr++
				key++
			}
			if max < curr {
				max = curr
			}
		}
	}
	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
