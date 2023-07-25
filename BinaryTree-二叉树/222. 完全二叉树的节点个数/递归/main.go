package 递归

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/20
 * Time: 22:04
 * Description: https://leetcode.cn/problems/count-complete-tree-nodes/
 */

type TreeNode = global.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 1
	if root.Left != nil {
		result += countNodes(root.Left)
	}
	if root.Right != nil {
		result += countNodes(root.Right)
	}
	return result
}
