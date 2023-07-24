package 前序遍历

import (
	"LC/BinaryTree-二叉树/global"
	"math"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/19
 * Time: 0:43
 * Description: https://leetcode.cn/problems/minimum-depth-of-binary-tree/
 */

type TreeNode = global.TreeNode

func main() {

}

var ans int

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = math.MaxInt
	getDepth(root, 1)
	return ans
}

func getDepth(node *TreeNode, depth int) {
	if node.Left == nil && node.Right == nil {
		ans = min(depth, ans)
		return
	}
	if node.Left != nil {
		getDepth(node.Left, depth+1)
	}
	if node.Right != nil {
		getDepth(node.Right, depth+1)
	}
	return
}

func min(leftDepth int, rightDepth int) int {
	if leftDepth < rightDepth {
		return leftDepth
	}
	return rightDepth
}
