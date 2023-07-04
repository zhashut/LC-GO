package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/16
 * Time: 21:08
 * Description: https://leetcode.cn/problems/intersection-of-two-linked-lists/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 方法一：快慢指针
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

// 方法二：双指针
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}

// 方法三：哈希集合
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}

	for temp := headA; temp != nil; temp = temp.Next {
		vis[temp] = true
	}

	for temp := headB; temp != nil; temp = temp.Next {
		if vis[temp] {
			return temp
		}
	}

	return nil
}
