package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/15
 * Time: 21:00
 * Description: https://leetcode.cn/problems/swap-nodes-in-pairs/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 1}
	l2 := ListNode{Val: 2}
	l3 := ListNode{Val: 3}
	l4 := ListNode{Val: 4}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4

	swapPairs2(&l1)
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for pt := dummy; pt != nil && pt.Next != nil && pt.Next.Next != nil; {
		pt, pt.Next, pt.Next.Next, pt.Next.Next.Next = pt.Next, pt.Next.Next, pt.Next.Next.Next, pt.Next
	}
	return dummy.Next
}

// 迭代
func swapPairs1(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	temp := dummyHead

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

// 递归1
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// 递归2
func swapPairs3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	one := head
	two := one.Next
	three := two.Next

	two.Next = one
	one.Next = swapPairs2(three)

	return two

}
