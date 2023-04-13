package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/13
 * Time: 9:48
 * Description: No Description
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := MyLinkedList{}
	list.AddAtTail(1)
	res := list.Get(0)
	fmt.Println(res)
}

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = this.Head
	this.Head = newNode
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{Val: val}

	if this.Head == nil {
		this.Head = newNode
	} else {
		cur := this.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}

	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}

	newNode := &ListNode{Val: val}
	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	cur.Next = newNode

	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
	} else {
		cur := this.Head
		for i := 0; i < index-1; i++ {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}

	this.Size--
}
