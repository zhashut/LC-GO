package main

import "container/list"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/12
 * Time: 1:11
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
	if root == nil {
		return nil
	}
	var stack = list.New() //栈
	res := []int{}         //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)     //弹出元素
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右左中
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
	}
	return res
}

// 中序遍历  右中左
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}         // 结果集
	var stack = list.New() // 栈
	stack.PushBack(root)
	var node *TreeNode

	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil { // 如果为空，则表示需要处理的是中间节点
			e = stack.Back() // 弹出元素（即中间节点）
			stack.Remove(e)  // 删除中间节点
			node = e.Value.(*TreeNode)
			ans = append(ans, node.Val) // 将中间节点加入到集合中
			continue                    // 继续弹出栈中下一个节点
		}
		// 压栈顺序： 右中左
		node = e.Value.(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		stack.PushBack(node) // 中间节点压入栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ans
}

// 后序遍历  中右左
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ans := []int{}         // 结果集
	var stack = list.New() // 栈
	stack.PushBack(root)
	var node *TreeNode

	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil { // 如果为空，则表示需要处理的是中间节点
			e = stack.Back() // 弹出元素（即中间节点）
			stack.Remove(e)  // 删除中间节点
			node = e.Value.(*TreeNode)
			ans = append(ans, node.Val) // 将中间节点加入到集合中
			continue                    // 继续弹出栈中下一个节点
		}
		// 压栈顺序： 右中左
		node = e.Value.(*TreeNode)
		stack.PushBack(node) // 中间节点压入栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return ans
}
