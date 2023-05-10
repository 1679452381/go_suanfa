/**
 * @author X
 * @date 2023/5/10
 * @description
 **/
package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	subsetsBack(&res, path, nums, 0)
	return res
}

func subsetsBack(res *[][]int, path []int, nums []int, start int) {
	if start > len(nums) {
		return
	}
	temp := make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		subsetsBack(res, path, nums, i+1)
		path = path[:len(path)-1]
	}
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
