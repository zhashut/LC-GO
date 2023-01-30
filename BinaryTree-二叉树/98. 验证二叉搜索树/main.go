package _8__验证二叉搜索树

import (
	"LC/BinaryTree-二叉树/global"
	"math"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/30
 * Time: 23:53
 * Description: No Description
 */

type TreeNode = global.TreeNode

/**
 有效的二叉搜索树定义如下：
	节点的左子树只包含 小于 当前节点的数。
	节点的右子树只包含 大于 当前节点的数。
	所有左子树和右子树自身必须也是二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root, math.MinInt64, math.MaxInt64)
}

func check(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}

	return check(node.Left, min, int64(node.Val)) && check(node.Right, int64(node.Val), max)
}

// 中序遍历解法
func isValidBST1(root *TreeNode) bool {
	// 保存上一个指针
	var prev *TreeNode
	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := travel(node.Left)
		// 当前值小于等于前一个节点的值，返回false
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}
	return travel(root)
}
