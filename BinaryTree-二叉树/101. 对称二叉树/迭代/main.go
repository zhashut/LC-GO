package main

import (
	"LC/BinaryTree-二叉树/global"
	"container/list"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/15
 * Time: 21:22
 * Description: No Description
 */

type TreeNode = global.TreeNode

// 迭代-队列
func isSymmetric_queue(root *TreeNode) bool {
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root.Left, root.Right)
	}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}

// 迭代-栈
func isSymmetric_stack(root *TreeNode) bool {
	stack := list.New()
	if root == nil {
		return true
	}
	stack.PushBack(root.Left)
	stack.PushBack(root.Right)

	for stack.Len() > 0 {
		right := stack.Remove(stack.Back()).(*TreeNode)
		left := stack.Remove(stack.Back()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		stack.PushBack(left.Left)
		stack.PushBack(right.Right)
		stack.PushBack(left.Right)
		stack.PushBack(right.Left)
	}
	return true
}
