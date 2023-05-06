package main

import (
	list2 "container/list"
)

func feinonaqi(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return feinonaqi(n-1) + feinonaqi(n-2)
}

func feinonaqi2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := []int{0, 1}
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

func feinonaqi3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	stack := list2.New()
	stack.PushBack(0)
	stack.PushBack(1)
	for i := 2; i <= n; i++ {
		a := stack.Back()
		stack.Remove(a)
		b := stack.Back()
		stack.Remove(b)

		stack.PushBack(a.Value.(int))
		stack.PushBack(a.Value.(int) + b.Value.(int))
	}
	a := stack.Back()
	res := stack.Remove(a)
	return res.(int)
}

//func main() {
//
//	fmt.Println(feinonaqi(6))
//	fmt.Println(feinonaqi2(6))
//	fmt.Println(feinonaqi3(6))
//	//fmt.Println(stack)
//}

//0 1 1 2 3 5
