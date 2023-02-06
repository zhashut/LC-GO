package 迭代

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/5
 * Time: 20:36
 * Description: No Description
 */

type TreeNode = global.TreeNode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	var pnode *TreeNode
	for node != nil {
		if val > node.Val {
			pnode = node
			node = node.Right
		} else {
			pnode = node
			node = node.Left
		}
	}
	if val > pnode.Val {
		pnode.Right = &TreeNode{Val: val}
	} else {
		pnode.Left = &TreeNode{Val: val}
	}
	return root
}
