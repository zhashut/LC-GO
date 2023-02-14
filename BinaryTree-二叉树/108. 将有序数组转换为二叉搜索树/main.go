package _08__将有序数组转换为二叉搜索树

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/2/14
 * Time: 22:07
 * Description: No Description
 */

type TreeNode = global.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) >> 1
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root

	// 方法二
	//root := &TreeNode{}
	//return travel(root, nums, 0, len(nums)-1)
}

// 方法二
func travel(root *TreeNode, nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) >> 1

	root = &TreeNode{Val: nums[mid]}
	root.Left = travel(root.Left, nums, left, mid-1)
	root.Right = travel(root.Right, nums, mid+1, right)

	return root
}
