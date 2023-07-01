package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/14
 * Time: 10:34
 * Description: https://leetcode.cn/problems/remove-linked-list-elements/
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
	l5 := ListNode{Val: 5}
	l6 := ListNode{Val: 6}

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5
	l5.Next = &l6

	removeElements(&l1, 2)
}

// 迭代
func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{Next: head}
	cur := prev
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return prev.Next
}

// 递归
func removeElements1(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
