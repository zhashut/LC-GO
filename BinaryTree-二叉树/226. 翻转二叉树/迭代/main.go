package main

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/14
 * Time: 22:07
 * Description: No Description
 */

type TreeNode = global.TreeNode

func main() {

}

// 迭代版本的前序遍历
func invertTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node.Left, node.Right = node.Right, node.Left // 中
			stack = append(stack, node)
			node = node.Left // 左
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right // 右
	}
	return root
}

// 迭代版本的后序遍历
func invertTree1(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	node := root
	var prev *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right == nil || node.Right == prev {
			node.Left, node.Right = node.Right, node.Left //交换
			prev = node
			node = nil
		} else {
			stack = append(stack, node)
			node = node.Right
		}
	}
	return root
}
