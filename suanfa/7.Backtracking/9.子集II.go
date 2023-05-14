/**
 * @author X
 * @date 2023/5/10
 * @description
 **/
package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	subsetsWithDupBack(&res, path, nums, 0)
	return res
}
func subsetsWithDupBack(res *[][]int, path []int, nums []int, start int) {
	//fmt.Println(path)
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		subsetsWithDupBack(res, path, nums, i+1)
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
