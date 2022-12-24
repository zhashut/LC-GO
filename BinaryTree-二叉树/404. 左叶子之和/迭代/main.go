package 迭代

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/24
 * Time: 22:42
 * Description: No Description
 */

type TreeNode = global.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	st := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	st = append(st, root)
	result := 0

	for len(st) != 0 {
		node := st[len(st)-1]
		st = st[:len(st)-1]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}
		if node.Right != nil {
			st = append(st, node.Right)
		}
		if node.Left != nil {
			st = append(st, node.Left)
		}
	}

	return result
}
