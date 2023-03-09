package _68__监控二叉树

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/9
 * Time: 21:03
 * Description: https://leetcode.cn/problems/binary-tree-cameras/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	result := 0

	var travel func(cur *TreeNode) int
	travel = func(cur *TreeNode) int {
		// 空节点，该节点有覆盖
		if cur == nil {
			return 2
		}

		// 后序遍历-左右中
		left := travel(cur.Left)
		right := travel(cur.Right)

		// 情况1
		// 左右节点都有覆盖
		if left == 2 && right == 2 {
			return 0
		}

		// 情况2
		// left == 0 && right == 0 左右节点无覆盖
		// left == 1 && right == 0 左节点有摄像头，右节点无覆盖
		// left == 0 && right == 1 左节点有无覆盖，右节点摄像头
		// left == 0 && right == 2 左节点无覆盖，右节点覆盖
		// left == 2 && right == 0 左节点覆盖，右节点无覆盖
		if left == 0 || right == 0 {
			result++
			return 1
		}

		// 情况3
		// left == 1 && right == 2 左节点有摄像头，右节点有覆盖
		// left == 2 && right == 1 左节点有覆盖，右节点有摄像头
		// left == 1 && right == 1 左右节点都有摄像头
		// 其他情况前段代码均已覆盖
		if left == 1 || right == 1 {
			return 2
		}

		// 以上代码我没有使用else，主要是为了把各个分支条件展现出来，这样代码有助于读者理解
		// 这个 return -1 逻辑不会走到这里。
		return -1
	}

	if travel(root) == 0 {
		result++
	}

	return result
}
