package main

import (
	"fmt"
	"math"
)

// 暴力解法 超时
func canCompleteCircuit1(gas []int, cost []int) int {
	for i := 0; i < len(cost); i++ {
		ret := gas[i] - cost[i]
		index := (i + 1) % len(cost)
		for ret > 0 && index != i {
			ret += gas[index] - cost[index]
			index = (index + 1) % len(cost)
		}
		if ret >= 0 && index == i {
			return i
		}
	}
	return -1
}

// 贪心一
// 思想：
// 情况一：如果gas的总和小于cost总和，那么无论从哪里出发，一定是跑不了一圈的
// 情况二：rest[i] = gas[i]-cost[i]为一天剩下的油，i从0开始计算累加到最后一站，如果累加没有出现负数，说明从0出发，油就没有断过，那么0就是起点。
// 情况三：如果累加的最小值是负数，汽车就要从非0节点出发，从后向前，看哪个节点能把这个负数填平，能把这个负数填平的节点就是出发节点。
func canCompleteCircuit2(gas []int, cost []int) int {
	sum := 0
	min := math.MaxInt
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if min > sum {
			min = sum
		}
	}
	if sum < 0 {
		return -1
	}
	if min >= 0 {
		return 0
	}
	for i := len(gas) - 1; i >= 0; i-- {
		rest := gas[i] - cost[i]
		min += rest
		if min >= 0 {
			return i
		}
	}
	return -1
}

func canCompleteCircuit3(gas []int, cost []int) int {
	start := 0
	sum := 0
	curSum := 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		curSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	if sum < 0 {
		return -1
	}
	return start
}
func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit1(gas, cost))
	fmt.Println(canCompleteCircuit2(gas, cost))
	fmt.Println(canCompleteCircuit3(gas, cost))
}
