package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/16
 * Time: 18:03
 * Description: No Description
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 栈 -----------------------------------------------------------------------
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{Val: 0, Next: head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// -----------------------------------------------------------------------

// 计算链表长度 -----------------------------------------------------------------------
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

// -----------------------------------------------------------------------

// 双指针 -----------------------------------------------------------------------
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

// -----------------------------------------------------------------------
