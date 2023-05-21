package main

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees1(root1.Left, root2.Left)
	root1.Right = mergeTrees1(root1.Right, root2.Right)
	return root1
}

func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := make([]*TreeNode, 0)

	queue = append(queue, root1, root2)
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i += 2 {
			node1 := queue[0]
			node2 := queue[1]
			queue = queue[2:]
			node1.Val += node2.Val
			if node1.Left != nil && node2.Left != nil {
				queue = append(queue, node1.Left, node2.Left)
			}
			if node1.Right != nil && node2.Right != nil {
				queue = append(queue, node1.Right, node2.Right)
			}
			if node1.Left == nil && node2.Left != nil {
				node1.Left = node2.Left
			}
			if node1.Right == nil && node2.Right != nil {
				node1.Right = node2.Right
			}
		}
	}
	return root1
}
