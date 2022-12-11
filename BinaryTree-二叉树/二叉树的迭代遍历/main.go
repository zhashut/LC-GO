package main

import "container/list"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/11
 * Time: 9:53
 * Description: No Description
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	if st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val) // 中
		if node.Right != nil {      //
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	st := list.New()
	cur := root

	for cur != nil || st.Len() > 0 {
		if cur != nil { // 指针来访问节点，访问到最底层
			st.PushBack(cur) // 将访问的节点放进栈
			cur = cur.Left   // 左
		} else {
			cur = st.Remove(st.Back()).(*TreeNode) // 从栈里弹出的数据，就是要处理的数据（放进ans数组里的数据）
			ans = append(ans, cur.Val)             // 中
			cur = cur.Right                        // 右
		}
	}
	return ans
}

// 后续遍历
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	return ans
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l, r = l+1, r-1
	}
}
