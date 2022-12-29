package _12__路径总和

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/29
 * Time: 20:09
 * Description: No Description
 */

type TreeNode = global.TreeNode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return traversal(root, targetSum-root.Val)

	// 精简版
	//targetSum -= root.Val // 将targetSum在遍历每层的时候都减去本层节点的值
	//if root.Left == nil && root.Right == nil && targetSum == 0 { // 如果剩余的targetSum为0, 则正好就是符合的结果
	//	return true
	//}
	//return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归找
}

func traversal(cur *TreeNode, num int) bool {
	if num == 0 && cur.Left == nil && cur.Right == nil { // 遇到叶子节点，并且计数为0
		return true
	}

	if cur.Left == nil && cur.Right == nil { // 遇到叶子节点直接返回
		return false
	}

	if cur.Left != nil { // 左
		num -= cur.Left.Val // 递归，处理节点;
		if traversal(cur.Left, num) {
			return true
		}
		num += cur.Left.Val // 回溯，撤销处理结果
	}

	if cur.Right != nil { // 右
		num -= cur.Right.Val // 递归，处理节点;
		if traversal(cur.Right, num) {
			return true
		}
		num += cur.Right.Val // 回溯，撤销处理结果
	}

	return false
}
