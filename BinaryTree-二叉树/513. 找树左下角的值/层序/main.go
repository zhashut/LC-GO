package 层序

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/25
 * Time: 20:51
 * Description: https://leetcode.cn/problems/find-bottom-left-tree-value/
 */

type TreeNode = global.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	result := 0
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				result = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return result
}
