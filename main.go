package main

import (
	"fmt"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/10/22
 * Time: 23:15
 * Description: No Description
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	/*	a1 := &ListNode{Val: 2}
		a2 := &ListNode{Val: 4}
		a3 := &ListNode{Val: 3}

		a1.Next = a2
		a2.Next = a3

		b1 := &ListNode{Val: 5}
		b2 := &ListNode{Val: 6}
		b3 := &ListNode{Val: 4}

		b1.Next = b2
		b2.Next = b3
	*/

	s := "abcabcbb"
	for i, _ := range s {
		fmt.Printf("i = %d, val = %v\n", i, s[i])
	}

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	ml, r, l := 0, 0, -1
	for l < len(s) {
		// 向前走,等于0说明没有出现过重复情况
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			freq[s[r+1]]++
			r++
		} else {
			freq[s[l]]--
			l++
		}
		ml = max(ml, r-l+1)
	}
	return ml
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func detectCycle(head *ListNode) *ListNode {
	// fast slow 的两个指针同时从head出发
	slow, fast := head, head
	// fast 第一轮每步为2跨步， slow 为1跨步 (第一轮是证明是否有环）
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 第二轮，说明有环，这时 fast 再次回到head，两指针每步走1步（找出链表入口点）
	fast = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// TODO map输出的顺序是随机的，题目要求升序！
func deleteDuplicates1(head *ListNode) (res *ListNode) {
	temp := make(map[int]int, 10)
	for head != nil {
		if _, ok := temp[head.Val]; ok {
			temp[head.Val] = -1
			head = head.Next
			continue
		}
		temp[head.Val] = head.Val
		head = head.Next
	}
	var tail *ListNode
	for _, val := range temp {
		if val != -1 {
			if res == nil {
				res = &ListNode{Val: val}
				tail = res
			} else {
				tail.Next = &ListNode{Val: val}
				tail = tail.Next
			}
		}
	}
	return
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func mergeTwoLists(l1, l2 *ListNode) (head *ListNode) {
	cur := &ListNode{Val: 0}
	head = cur
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}

	return head.Next
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
