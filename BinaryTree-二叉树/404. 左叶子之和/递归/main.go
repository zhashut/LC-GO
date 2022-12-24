package 递归

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/24
 * Time: 22:13
 * Description: No Description
 */

type TreeNode = global.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}

	leftVal := sumOfLeftLeaves(root.Left)                                    // 左
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // 左子树就是一个左叶子的情况
		leftVal += root.Left.Val
	}

	rightVal := sumOfLeftLeaves(root.Right) // 右

	sum := leftVal + rightVal // 中
	return sum
}

// 精简版
func sumOfLeftLeaves1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := sumOfLeftLeaves(root.Left) // 左

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue = root.Left.Val // 中
	}

	rightValue := sumOfLeftLeaves(root.Right) // 右

	return leftValue + rightValue
}
