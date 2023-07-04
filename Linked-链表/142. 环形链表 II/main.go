package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/17
 * Time: 8:11
 * Description: https://leetcode.cn/problems/linked-list-cycle-ii/
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 方法一：哈希表
func detectCycle(head *ListNode) *ListNode {
	ans := map[*ListNode]int{}
	for node := head; node != nil; node = node.Next {
		if _, ok := ans[node]; ok {
			return node
		}
		ans[node]++
	}

	return nil
}

// 快慢指针
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
