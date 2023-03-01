package main

import (
	"math"
	"sort"
)

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/28
 * Time: 21:05
 * Description: No Description
 */

func main() {
	nums, k := []int{2, -3, -1, 5, -4}, 2
	largestSumAfterKNegations(nums, k)
}

func largestSumAfterKNegations(nums []int, k int) int {
	// 绝对值比较大小排序
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	result := 0
	for i := 0; i < len(nums)-1; i++ {
		if k > 0 && nums[i] < 0 {
			result += -nums[i]
			k--
			continue
		}
		result += nums[i]
	}

	if k%2 == 1 {
		result += -nums[len(nums)-1]
	} else {
		result += nums[len(nums)-1]
	}

	return result
}
