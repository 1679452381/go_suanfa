/**
 * @author X
 * @date 2023/6/15
 * @description
 **/
package main

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}
	res := make([]int, 0)
	//	基准值
	pivot := nums[0]
	minNums := make([]int, 0)
	maxNums := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			minNums = append(minNums, nums[i])
		} else {
			maxNums = append(maxNums, nums[i])
		}
	}
	res = append(res, quickSort(minNums)...)
	res = append(res, pivot)
	res = append(res, quickSort(maxNums)...)
	return res
}

func ArrSort(nums []int) []int {
	return quickSort2(nums, 0, len(nums)-1)
}
func quickSort2(nums []int, left, right int) []int {
	if left > right {
		return nil
	}
	i, j, pivot := left, right, nums[left]
	for i < j {
		// 寻找小于主元的右边元素
		for i < j && nums[j] >= pivot {
			j--
		}
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	quickSort2(nums, left, i-1)
	quickSort2(nums, i+1, right)
	return nums
}
func main() {
	fmt.Println(quickSort([]int{2, 3, 1, 5, 56, 82, 6, 9}))
	fmt.Println(ArrSort([]int{2, 3, 1, 5, 56, 82, 6, 9}))

}
