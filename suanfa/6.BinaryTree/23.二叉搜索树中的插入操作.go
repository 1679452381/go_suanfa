package main

// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 迭代
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	pre, cur := root, root
	for cur != nil {
		pre = cur
		if cur.Val > val {
			cur = cur.Left
		} else if cur.Val < val {
			cur = cur.Right
		}
	}
	node := &TreeNode{Val: val}
	if val < pre.Val {
		pre.Left = node
	} else if val > pre.Val {
		pre.Right = node
	}
	return root
}
