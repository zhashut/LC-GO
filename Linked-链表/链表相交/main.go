package 链表相交

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/15
 * Time: 20:54
 * Description: https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// 双指针
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = headB
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = headA
		}
	}

	return l1
}
