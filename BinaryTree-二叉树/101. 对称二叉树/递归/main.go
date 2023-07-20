package main

import (
	"LC/BinaryTree-二叉树/global"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/15
 * Time: 21:17
 * Description: https://leetcode.cn/problems/symmetric-tree/
 */

type TreeNode = global.TreeNode

func main() {

}

// 后序递归
func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root.Left, root.Right)
}
