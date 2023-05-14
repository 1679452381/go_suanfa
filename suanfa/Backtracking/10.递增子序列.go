package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	findSubsequencesBack(nums, &res, path, 0, used)
	return res
}

func findSubsequencesBack(nums []int, res *[][]int, path []int, start int, used []bool) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] != true {
			continue
		}
		if len(path) == 0 {
			path = append(path, nums[i])
			used[i] = true
		} else {
			if nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[i] = true
			} else {
				return
			}
		}
		findSubsequencesBack(nums, res, path, i+1, used)
		path = path[:len(path)-1]
		used[i] = false
	}

}
func main() {
	nums := []int{4, 6, 7, 7}
	res := findSubsequences(nums)
	fmt.Println(res)
}
