package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	ret := make([]int, 0)
	record := make(map[int]int)
	for _, val := range nums1 {
		record[val]++
	}
	for _, val := range nums2 {
		if record[val] != 0 {
			record[val] = 0
			ret = append(ret, val)
		}
	}
	return ret
}

func main() {
	s := []int{1, 2, 2, 1, 3}
	t := []int{2, 3, 3}
	fmt.Println(intersection(s, t))
}
