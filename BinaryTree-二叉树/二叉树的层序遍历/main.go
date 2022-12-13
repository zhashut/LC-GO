package main

import "container/list"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/14
 * Time: 1:34
 * Description: No Description
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	que := list.New()
	if root != nil {
		que.PushBack(root)
	}
	ans := [][]int{}
	for que.Len() > 0 {
		size := que.Len() // 保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		tmp := []int{}    // 临时数组，在这个位置有清空上一组的作用
		for i := 0; i < size; i++ {
			node := que.Remove(que.Front()).(*TreeNode) // 出队列
			tmp = append(tmp, node.Val)                 // 将值加入本层切片中
			if node.Left != nil {
				que.PushBack(node.Left)
			}
			if node.Right != nil {
				que.PushBack(node.Right)
			}
		}
		ans = append(ans, tmp) // 放入结果集
	}
	return ans
}

// 二叉树的递归遍历
func levelOrder1(root *TreeNode) [][]int {
	arr := [][]int{}
	depth := 0

	var order func(node *TreeNode, depth int)

	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}

		arr[depth] = append(arr[depth], node.Val)
		order(node.Left, depth+1)
		order(node.Right, depth+1)
	}
	order(root, depth)

	return arr
}
