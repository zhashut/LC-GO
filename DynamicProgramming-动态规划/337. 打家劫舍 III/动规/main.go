package 动规

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/27
 * Time: 19:00
 * Description: No Description
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func robTree(cur *TreeNode) []int {
	if cur == nil {
		return []int{0, 0}
	}
	// 后序遍历
	left := robTree(cur.Left)
	right := robTree(cur.Right)

	// 考虑去偷当前屋子
	robCur := cur.Val + left[0] + right[0]
	// 考虑不去偷当前屋子
	notRobCur := max(left[0], left[1]) + max(right[0], right[1])

	// 注意顺序：0 不偷， 1 去偷
	return []int{notRobCur, robCur}
}

// 解题二：用整数当变量
func rob337(root *TreeNode) int {
	a, b := dfsTreeRob(root)
	return max(a, b)
}

func dfsTreeRob(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	l0, l1 := dfsTreeRob(root.Left)
	r0, r1 := dfsTreeRob(root.Right)
	// 当前节点没有被打劫
	tmp0 := max(l0, l1) + max(r0, r1)
	// 当前节点被打劫
	tmp1 := root.Val + l0 + r0
	return tmp0, tmp1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
