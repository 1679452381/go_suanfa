package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	findSubsequencesBack(nums, &res, path, 0)
	return res
}

func findSubsequencesBack(nums []int, res *[][]int, path []int, start int) {
	if len(path) > 1 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		//if i > 0 && nums[i] == nums[i-1] && used[i-1] != true {
		//	continue
		//}
		if used[nums[i]] {
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			findSubsequencesBack(nums, res, path, i+1)
			path = path[:len(path)-1]
		}
	}

}
func main() {
	nums := []int{4, 6, 7, 7}
	res := findSubsequences(nums)
	fmt.Println(res)
}
