package main

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/14
 * Time: 21:46
 * Description: No Description
 */

type TreeNode = global.TreeNode

func main() {

}

// 递归-前序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left // 交换
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 递归-后序遍历
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)                         //遍历左节点
	invertTree(root.Right)                        //遍历右节点
	root.Left, root.Right = root.Right, root.Left //交换
	return root
}
