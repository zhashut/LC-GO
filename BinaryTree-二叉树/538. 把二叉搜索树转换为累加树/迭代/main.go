package 迭代

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/15
 * Time: 23:45
 * Description: No Description
 */

type TreeNode = global.TreeNode

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	cur := root
	stack := list.New()
	pre := 0

	for cur != nil || stack.Len() != 0 {
		for cur != nil {
			stack.PushBack(cur)
			cur = cur.Right
		}
		cur = stack.Remove(stack.Back()).(*TreeNode)
		cur.Val += pre
		pre = cur.Val
		cur = cur.Left
	}

	return root
}
