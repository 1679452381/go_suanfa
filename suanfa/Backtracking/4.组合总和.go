package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	combinationSumBack(&res, path, candidates, target, 0, 0)
	return res
}
func combinationSumBack(res *[][]int, path, candidates []int, target, sum, start int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		//剪枝优化
		if sum+candidates[i] > target {
			continue
		}
		path = append(path, candidates[i])
		sum += candidates[i]
		combinationSumBack(res, path, candidates, target, sum, i)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
