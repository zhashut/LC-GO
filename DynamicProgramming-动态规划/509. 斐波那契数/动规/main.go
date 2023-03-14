package 动规

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/14
 * Time: 21:41
 * Description: https://leetcode.cn/problems/fibonacci-number/
 */

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}

	return c
}
