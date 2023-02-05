package 迭代

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/4
 * Time: 20:49
 * Description: No Description
 */

type TreeNode = global.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
	}

	return root
}
