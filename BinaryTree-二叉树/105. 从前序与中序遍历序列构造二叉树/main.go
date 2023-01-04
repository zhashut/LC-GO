package _05__从前序与中序遍历序列构造二叉树

import "LC/BinaryTree-二叉树/global"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/4
 * Time: 20:31
 * Description: No Description
 */

type TreeNode = global.TreeNode

var (
	hash map[int]int
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	hash = make(map[int]int, len(inorder))
	for i, v := range inorder {
		hash[v] = i
	}
	root := build(preorder, inorder, 0, 0, len(inorder)-1)
	return root
}

func build(pre []int, in []int, root, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	rootVal := pre[root]
	index := hash[rootVal]
	node := &TreeNode{Val: rootVal}
	node.Left = build(pre, in, root+1, l, index-1)
	node.Right = build(pre, in, root+(index-l)+1, index+1, r)

	return node
}

//------------------------------------------------------------------------------------------------------------------------------------
// 解法二, 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存, 内存使用(leetcode test case) 4.7MB -> 4.3MB.
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
			root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
		}
	}
	return root
}

//------------------------------------------------------------------------------------------------------------------------------------

// 解法二
func buildTree2(preorder []int, inorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildPreIn2TreeDFS(preorder, 0, len(preorder)-1, 0, inPos)
}

func buildPreIn2TreeDFS(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: pre[preStart]}
	rootIdx := inPos[pre[preStart]]
	leftLen := rootIdx - inStart
	root.Left = buildPreIn2TreeDFS(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildPreIn2TreeDFS(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root
}

//------------------------------------------------------------------------------------------------------------------------------------
