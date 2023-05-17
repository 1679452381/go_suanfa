package main

// 反转二叉树
// 1.递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 2.深度优先遍历
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) > 0 {
		cur := s[len(s)-1]
		s = s[0 : len(s)-1]
		temp := cur.Left
		cur.Left = cur.Right
		cur.Right = temp
		if cur.Right != nil {
			s = append(s, cur.Right)
		}
		if cur.Left != nil {
			s = append(s, cur.Left)
		}
	}
	return root
}

// 3.广度优先遍历
func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			temp := cur.Left
			cur.Left = cur.Right
			cur.Right = temp
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
		}
	}
	return root
}
