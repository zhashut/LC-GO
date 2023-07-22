package main

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/17
 * Time: 12:17
 * Description: https://leetcode.cn/problems/maximum-depth-of-binary-tree/
 */

type TreeNode = global.TreeNode

func main() {

}

func maxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()
		depth++
		for i := 0; i < length; i++ {
			e := queue.Remove(queue.Front()).(*TreeNode)
			if e.Left != nil {
				queue.PushBack(e.Left)
			}
			if e.Right != nil {
				queue.PushBack(e.Right)
			}
		}
	}
	return depth
}
