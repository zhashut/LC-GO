package 利用完全二叉树特性的递归解法

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/20
 * Time: 22:20
 * Description: No Description
 */

type TreeNode = global.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 开始根据做深度和有深度是否相同来判断该子树是不是满二叉树
	leftH, rightH := 0, 0 // 这里初始为0是有目的的，为了下面求指数方便
	leftNode, rightNode := root.Left, root.Right
	for leftNode != nil { // 求左子树深度
		leftNode = leftNode.Left
		leftH++
	}

	for rightNode != nil { // 求右子树深度
		rightNode = rightNode.Right
		rightH++
	}

	if leftH == rightH {
		return (2 << leftH) - 1 // 注意(2<<1) 相当于2^2，返回满足满二叉树的子树节点数量
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
