package _01__二叉搜索树中的众数

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/2
 * Time: 19:50
 * Description: No Description
 */

type TreeNode = global.TreeNode

// 中序遍历
func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	count := 1
	max := 1
	var pre *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		pre = node
		travel(node.Right)
	}
	travel(root)
	return res
}
