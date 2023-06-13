/**
 * @author X
 * @date 2023/6/12
 * @description
 **/
package main

import (
	"fmt"
	"sort"
)

// 回溯 会超时
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0, 3)
	res := make([][]int, 0)
	sum := 0
	used := make([]int, len(nums))
	fmt.Println(nums)
	fmt.Println(used)
	threeSumBack(nums, path, used, &res, sum, 0)
	return res
}
func threeSumBack(nums, path, used []int, res *[][]int, sum, start int) {
	if len(path) == 3 {
		if sum == 0 {
			temp := make([]int, 3)
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}
	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
			continue
		}
		used[i] = 1
		sum += nums[i]
		path = append(path, nums[i])
		threeSumBack(nums, path, used, res, sum, i+1)
		used[i] = 0
		sum -= nums[i]
		path = path[:len(path)-1]
	}
}

// 双指针
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4, 0, -1, -3, -2, -2}

	fmt.Println(threeSum1(nums))
	fmt.Println(threeSum(nums))
}
