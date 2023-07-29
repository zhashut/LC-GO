package 递归

import (
	"LC/BinaryTree-二叉树/global"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/25
 * Time: 21:09
 * Description: https://leetcode.cn/problems/find-bottom-left-tree-value/
 */

type TreeNode = global.TreeNode

var depth int  // 全局变量 最大深度
var result int // 记录最终结果

func findBottomLeftValue(root *TreeNode) int {
	depth, result = 0, 0 // 初始化
	traversal(root, 1)
	return result
}

func traversal(root *TreeNode, d int) {
	if root == nil {
		return
	}
	// 因为先遍历左边，所以左边如果有值，右边的同层不会更新结果
	if root.Left == nil && root.Right == nil && depth < d {
		depth = d
		result = root.Val
	}
	traversal(root.Left, d+1)
	traversal(root.Right, d+1)
}
