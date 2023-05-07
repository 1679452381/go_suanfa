package main

import "fmt"

// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0, k)
	back(&ret, path, n, k, 1)
	return ret

}

func back(ret *[][]int, path []int, n, k, start int) {
	if len(path) == k {
		temp := make([]int, k)
		copy(temp, path)
		*ret = append(*ret, temp)
		return
	}
	//剪枝优化
	//已经选择的元素个数：path.size();
	//所需需要的元素个数为: k - path.size();
	//列表中剩余元素（n-i） >= 所需需要的元素个数（k - path.size()）
	//在集合n中至多要从该起始位置 : i <= n - (k - path.size()) + 1，开始遍历
	for i := start; i <= n-(k-len(path))+1; i++ {
		//for i := start; i <= n; i++ {
		path = append(path, i)
		back(ret, path, n, k, i+1)
		path = path[0 : len(path)-1]
	}
}

func main() {
	ret := combine(4, 2)
	fmt.Println(ret)
}
