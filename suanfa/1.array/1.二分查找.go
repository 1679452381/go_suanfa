package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + (end-start)/2
		//fmt.Println(start, end, mid)
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	a := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(a, 2))
}
