package main

import (
	"math"
	"sort"
)

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/10/30
 * Time: 3:04
 * Description: No Description
 */

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	threeSumClosest(nums, target)
}

func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				// TODO: 这里我理解为控制最接近 target 的数进入，diff是在不断变小的
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
