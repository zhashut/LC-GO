package main

import (
	"LC/BinaryTree-二叉树/global"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/12/31
 * Time: 20:44
 * Description: No Description
 */

type TreeNode = global.TreeNode

var (
	hash map[int]int
)

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}

// 解法一, 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存, 内存使用(leetcode test case) 4.7MB -> 4.3MB.
func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
			root.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return root
}

// ---------------------------------------------------解法二----------------------------------------------------------------------------
func buildTree1(inorder []int, postorder []int) *TreeNode {
	hash = make(map[int]int)
	for i, v := range inorder { // 用map保存中序序列的数值对应位置
		hash[v] = i
	}

	// 以左闭右闭的原则进行切分
	root := rebuild(inorder, postorder, len(postorder)-1, 0, len(inorder)-1)
	return root
}

// rootIdx表示根节点在后序数组中的索引，l, r 表示在中序数组中的前后切分点
func rebuild(inorder []int, postorder []int, rootIdx int, l, r int) *TreeNode {
	if l > r { // 说明没有元素，返回空树
		return nil
	}
	if l == r { // 只剩唯一一个元素，直接返回
		return &TreeNode{Val: inorder[l]}
	}

	rootV := postorder[rootIdx]   // 根据后序数组找到根节点的值
	rootIn := hash[rootV]         // 找到根节点在对应的中序数组中的位置
	root := &TreeNode{Val: rootV} // 构造根节点
	// 重建左节点和右节点
	root.Left = rebuild(inorder, postorder, rootIdx-(r-rootIn)-1, l, rootIn-1)
	root.Right = rebuild(inorder, postorder, rootIdx-1, rootIn+1, r)
	return root
}

//-------------------------------------------------------------------------------------------------------------------------------

//--------------------------------------------------解法三-----------------------------------------------------------------------------
func buildTree2(inorder []int, postorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildInPos2TreeDFS(postorder, 0, len(postorder)-1, 0, inPos)
}

func buildInPos2TreeDFS(post []int, postStart int, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{Val: post[postEnd]}
	rootIdx := inPos[post[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = buildInPos2TreeDFS(post, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildInPos2TreeDFS(post, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}

//-------------------------------------------------------------------------------------------------------------------------------
