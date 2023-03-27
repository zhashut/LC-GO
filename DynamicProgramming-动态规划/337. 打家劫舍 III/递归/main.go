package 递归

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/27
 * Time: 18:28
 * Description: https://leetcode.cn/problems/house-robber-iii/
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	umap = make(map[*TreeNode]int, 0)
)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	if umap[root] != 0 {
		return umap[root]
	}
	// 偷父节点
	rootV := root.Val
	if root.Left != nil {
		rootV += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		rootV += rob(root.Right.Left) + rob(root.Right.Right)
	}
	// 不偷父节点
	nodeV := rob(root.Left) + rob(root.Right)
	umap[root] = max(rootV, nodeV)

	return max(rootV, nodeV)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
