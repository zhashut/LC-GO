package 递归

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/14
 * Time: 21:54
 * Description: https://leetcode.cn/problems/fibonacci-number/
 */

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
