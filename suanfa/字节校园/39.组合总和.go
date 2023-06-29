/**
 * @author X
 * @date 2023/6/26
 * @description
 **/
package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	combinationSumBack(candidates, path, target, 0, 0, &res)
	return res
}
func combinationSumBack(candidates, path []int, target, start, sum int, res *[][]int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(candidates) && sum < target; i++ {
		//if target == candidates[i] {
		//	break
		//}
		path = append(path, candidates[i])
		//target -= candidates[i]
		sum += candidates[i]
		combinationSumBack(candidates, path, target, i, sum, res)
		path = path[:len(path)-1]
		//target += candidates[i]
		sum -= candidates[i]
	}
}

func main() {

	candidates := []int{2, 2, 3, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
