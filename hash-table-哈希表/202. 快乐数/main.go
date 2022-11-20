package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/20
 * Time: 0:22
 * Description: No Description
 */

func main() {
	isHappy1(4)
}

// 用哈希集合检测循环
func isHappy(n int) bool {
	m := map[int]bool{}
	for ; n != 1 && !m[n]; n, m[n] = step(n), true {
	}
	return n == 1
}

// 慢指针法
func isHappy1(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) (sum int) {
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return
}
