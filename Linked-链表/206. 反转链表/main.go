package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/15
 * Time: 20:00
 * Description: https://leetcode.cn/problems/reverse-linked-list/
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

	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	l4.Next = &l5

	reverseList1(&l1)
}

// 迭代
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// 递归
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	// 使用指针初始化时，初始化值为 nil
	var ans *ListNode

	for x := head; x != nil; x = x.Next {
		ans = &ListNode{
			Val:  x.Val,
			Next: ans,
		}
	}

	return ans
}
