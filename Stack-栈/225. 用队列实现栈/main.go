package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/30
 * Time: 21:14
 * Description: No Description
 */

func main() {

}

//  -------------- 方法一,使用两个队列-----------------------------

//type MyStack struct {
//	Queue1 []int
//	Queue2 []int
//}
//
//func Constructor() MyStack {
//	return MyStack{
//		Queue1: make([]int, 0),
//		Queue2: make([]int, 0),
//	}
//}
//
//func (this *MyStack) Push(x int) {
//	this.Queue2 = append(this.Queue2, x)
//	this.Move()
//}
//
//func (this *MyStack) Move() {
//	if len(this.Queue1) == 0 {
//		this.Queue1, this.Queue2 = this.Queue2, this.Queue1
//	} else {
//		this.Queue2 = append(this.Queue2, this.Queue1[0])
//		this.Queue1 = this.Queue1[1:]
//		this.Move()
//	}
//}
//
//func (this *MyStack) Pop() int {
//	val := this.Queue1[0]
//	this.Queue1 = this.Queue1[1:]
//	return val
//}
//
//func (this *MyStack) Top() int {
//	return this.Queue1[0]
//}
//
//func (this *MyStack) Empty() bool {
//	return len(this.Queue1) == 0
//}

//  -------------- 方法一,使用两个队列-----------------------------

//  -------------- 方法二,使用一个队列-----------------------------

type MyStack struct {
	Queue []int
}

func Constructor() MyStack {
	return MyStack{
		Queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
	n := len(this.Queue) - 1
	for n != 0 {
		val := this.Queue[0]
		this.Queue = this.Queue[1:]
		this.Queue = append(this.Queue, val)
		n--
	}
	val := this.Queue[0]
	this.Queue = this.Queue[1:]
	return val
}

func (this *MyStack) Top() int {
	val := this.Pop()
	this.Queue = append(this.Queue, val)
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.Queue) == 0
}

//  -------------- 方法二,使用一个队列-----------------------------

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
