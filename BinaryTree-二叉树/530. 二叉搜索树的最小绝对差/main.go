package _30__二叉搜索树的最小绝对差

import (
	"LC/BinaryTree-二叉树/global"
	"math"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/31
 * Time: 22:17
 * Description: No Description
 */

type TreeNode = global.TreeNode

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	min := math.MaxInt64
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return min
}
