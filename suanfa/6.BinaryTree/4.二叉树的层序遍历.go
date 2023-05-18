package main

import "fmt"

// 层序遍历
// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	for len(queue) > 0 {
		nextQueue := make([]*TreeNode, 0)
		path := make([]int, 0)
		n := len(queue)
		for i := 0; i < n; i++ {
			path = append(path, queue[i].Val)
			if queue[i].Left != nil {
				nextQueue = append(nextQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				nextQueue = append(nextQueue, queue[i].Right)
			}
		}
		queue = nextQueue
		res = append(res, path)
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)

	queue = append(queue, root)
	for len(queue) > 0 {
		path := make([]int, 0)
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			path = append(path, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, path)
	}
	return res
}

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(levelOrder(root))
}
