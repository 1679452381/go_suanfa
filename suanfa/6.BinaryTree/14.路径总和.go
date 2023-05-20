package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	qNode := make([]*TreeNode, 0)
	qSum := make([]int, 0)
	qNode = append(qNode, root)
	qSum = append(qSum, root.Val)
	for len(qNode) > 0 {
		n := len(qNode)
		for i := 0; i < n; i++ {
			cur := qNode[0]
			qNode = qNode[1:]
			sum := qSum[0]
			qSum = qSum[1:]
			if cur.Right == nil && cur.Left == nil {
				if sum == targetSum {
					return true
				}
				continue
			}

			if cur.Left != nil {
				qNode = append(qNode, cur.Left)
				qSum = append(qSum, sum+cur.Left.Val)
			}
			if cur.Right != nil {
				qNode = append(qNode, cur.Right)
				qSum = append(qSum, sum+cur.Right.Val)
			}
		}
	}
	return false
}
