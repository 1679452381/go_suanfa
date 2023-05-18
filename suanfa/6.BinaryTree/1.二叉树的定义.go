/**
 * @author X
 * @date 2023/5/9
 * @description
 **/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 2.

// 二叉树的最大深度
// 二叉树的最小深度
// 完全二叉树的节点个数
// 平衡二叉树
// 二叉树的所有路径
// 左叶子之和
// 找树左下角的值
// 路经总和
// 利用中序和后序遍历构建二叉树
// 最大二叉树
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
	fmt.Println(levelOrder2(root))
}
