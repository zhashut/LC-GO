package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/16
 * Time: 21:08
 * Description: No Description
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 方法一：哈希集合
func getIntersectionNode(headA, headB *ListNode) *ListNode {
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
