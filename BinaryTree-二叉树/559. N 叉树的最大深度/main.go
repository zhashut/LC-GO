package _59__N_叉树的最大深度

import "container/list"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/17
 * Time: 13:10
 * Description: https://leetcode.cn/problems/maximum-depth-of-binary-tree/
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for j := range node.Children {
				queue.PushBack(node.Children[j])
			}
		}
		depth++
	}
	return depth
}
