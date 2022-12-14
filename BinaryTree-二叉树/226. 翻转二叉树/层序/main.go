package main

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/14
 * Time: 22:25
 * Description: No Description
 */

type TreeNode = global.TreeNode

func main() {

}

// 层序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	node := root
	queue.PushBack(node)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			e.Left, e.Right = e.Right, e.Left
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return root
}
