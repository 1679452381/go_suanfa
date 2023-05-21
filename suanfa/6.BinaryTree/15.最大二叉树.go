package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	maxVal := 0
	maxValIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxValIndex = i
		}
	}
	root := &TreeNode{Val: maxVal}
	if 0 < maxValIndex {
		root.Left = constructMaximumBinaryTree(nums[0:maxValIndex])
	}
	if maxValIndex+1 < len(nums) {
		root.Right = constructMaximumBinaryTree(nums[maxValIndex+1 : len(nums)])
	}

	return root
}
