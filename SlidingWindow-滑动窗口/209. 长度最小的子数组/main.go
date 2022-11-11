package main

import (
	"math"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/10
 * Time: 21:53
 * Description: No Description
 */

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	minSubArrayLen(target, nums)
}

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	start, end, ans := 0, 0, math.MaxInt32
	sum := 0
	for end < n {
		sum += nums[end]
		if sum >= s {
			ans = min(ans, end-start-1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
