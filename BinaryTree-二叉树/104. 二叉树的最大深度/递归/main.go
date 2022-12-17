package main

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/17
 * Time: 13:01
 * Description: No Description
 */

type TreeNode = global.TreeNode

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxdepth(root.Left), maxdepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 遍历
func maxdepth(root *TreeNode) int {
	levl := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}
