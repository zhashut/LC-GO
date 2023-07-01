package main

import (
	"math"
	"sort"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/10
 * Time: 21:53
 * Description: https://leetcode.cn/problems/minimum-size-subarray-sum/
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
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
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

func minSubArrayLen1(s int, nums []int) int {
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
