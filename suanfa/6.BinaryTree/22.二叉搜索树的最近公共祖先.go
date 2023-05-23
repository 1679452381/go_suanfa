package main

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestor(root.Left, p, q)
		if left != nil {
			return left
		}
	}

	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestor(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	return root
}
