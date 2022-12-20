package main

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/20
 * Time: 21:44
 * Description: No Description
 */

type TreeNode = global.TreeNode

func main() {

}

// 层序遍历
func countNodes(root *TreeNode) int {
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	result := 0
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			result++
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
