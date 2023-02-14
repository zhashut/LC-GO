package 迭代

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/14
 * Time: 21:59
 * Description: No Description
 */

type TreeNode = global.TreeNode

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	// 处理 root，让 root 移动到[low, high] 范围内，注意是左闭右闭
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	cur := root
	// 此时 root 已经在[low, high] 范围内，处理左孩子元素小于 low 的情况（左节点是一定小于 root.Val，因此天然小于 high）
	for cur != nil {
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}

	cur = root
	// 此时 root 已经在[low, high] 范围内，处理右孩子大于 high 的情况
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}

	return root
}
