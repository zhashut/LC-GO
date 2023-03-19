package _43__整数拆分

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/19
 * Time: 18:13
 * Description: https://leetcode.cn/problems/integer-break/
 */

/*
*
动态五部曲
1.确定dp下标及其含义
2.确定递推公式
3.确定dp初始化
4.确定遍历顺序
5.打印dp
*
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// i可以差分为i-j和j。由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，
			//在求最大值的时候，一个是j与i-j相乘，一个是j与dp[i-j].
			dp[i] = max(j*(i-j), j*dp[i-j], dp[i])
		}
	}

	return dp[n]
}

func max(params ...int) int {
	max := params[0]
	if max < params[1] {
		max = params[1]
	}

	if max < params[2] {
		max = params[2]
	}

	return max
}
