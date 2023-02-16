package 递归

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/15
 * Time: 23:05
 * Description: No Description
 */

type TreeNode = global.TreeNode

// 倒序递归-右中左
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	pre := 0
	travel(root, &pre)

	return root
}

func travel(cur *TreeNode, pre *int) {
	if cur == nil {
		return
	}

	travel(cur.Right, pre)
	cur.Val += *pre
	*pre = cur.Val
	travel(cur.Left, pre)
}
