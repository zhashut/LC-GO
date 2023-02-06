package 递归

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/5
 * Time: 20:22
 * Description: No Description
 */

type TreeNode = global.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
