package main

// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字1到9
// 每个数字 最多使用一次

import "fmt"

func combinationSum3(k int, n int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0, k)
	back(&ret, path, n, k, 1, 0)
	return ret
}
func back(ret *[][]int, path []int, n, k, start, sum int) {
	if len(path) == k {
		if sum == n {
			temp := make([]int, k)
			copy(temp, path)
			*ret = append(*ret, temp)
		}
		return
	}
	//剪枝优化
	for i := start; i <= 9-(k-len(path))+1; i++ {
		sum += i
		path = append(path, i)
		back(ret, path, n, k, i+1, sum)
		sum -= i
		path = path[:len(path)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
