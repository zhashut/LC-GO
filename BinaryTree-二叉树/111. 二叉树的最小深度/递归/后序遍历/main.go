package main

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/19
 * Time: 0:09
 * Description: https://leetcode.cn/problems/minimum-depth-of-binary-tree/
 */

type TreeNode = global.TreeNode

func main() {

}

// 后序递归
func minDepth(root *TreeNode) int {
	return getDepth(root)

	// 精简版
	//if root == nil {
	//	return 0
	//}
	//if root.Left == nil && root.Right != nil {
	//	return 1 + minDepth(root.Right)
	//}
	//if root.Left != nil && root.Right == nil {
	//	return 1 + minDepth(root.Left)
	//}
	//return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := getDepth(node.Left)   // 左
	rightDepth := getDepth(node.Right) // 右

	if node.Left == nil && node.Right != nil { // 中
		return 1 + rightDepth
	}
	if node.Left != nil && node.Right == nil {
		return 1 + leftDepth
	}

	result := 1 + min(leftDepth, rightDepth)
	return result
}

func min(leftDepth int, rightDepth int) int {
	if leftDepth < rightDepth {
		return leftDepth
	}
	return rightDepth
}
