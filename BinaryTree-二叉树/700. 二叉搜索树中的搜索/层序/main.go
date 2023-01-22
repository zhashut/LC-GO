package 层序

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/22
 * Time: 20:04
 * Description: No Description
 */

type TreeNode = global.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		len := queue.Len()
		for i := 0; i < len; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			if e.Left != nil {
				if e.Left.Val == val {
					return e.Left
				}
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				if e.Right.Val == val {
					return e.Right
				}
				queue.PushBack(e.Right)
			}
		}
	}
	return nil
}
