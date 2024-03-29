package main

import "fmt"

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/6/30
 * Time: 7:43
 * Description: https://leetcode.cn/problems/binary-search/
 */

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	index := search(nums, target)
	fmt.Println(index)
}

// 左闭右开
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}

// 左闭右闭
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
