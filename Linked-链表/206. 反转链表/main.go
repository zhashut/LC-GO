package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/15
 * Time: 20:00
 * Description: No Description
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
	if head == nil {
		return &ListNode{}
	}

	var prev *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
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

// TODO: 这里因为初始化是0，不是nil所以反转后会多一个0，因此通不过， java可以初始化为 nil
func reverseList2(head *ListNode) *ListNode {
	ans := &ListNode{}

	for x := head; x != nil; x = x.Next {
		ans = &ListNode{
			Val:  x.Val,
			Next: ans,
		}
	}

	return ans
}
