package _10__平衡二叉树

import (
	"LC/BinaryTree-二叉树/global"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/29
 * Time: 17:28
 * Description: https://leetcode.cn/problems/balanced-binary-tree/
 */

type TreeNode = global.TreeNode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHight := depth(root.Left)
	rightHight := depth(root.Right)
	return abs(leftHight-rightHight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
