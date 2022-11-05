package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/6
 * Time: 0:54
 * Description: No Description
 */

func main() {
	nums := []int{1, 3, 5, 6}
	res := searchInsert1(nums, 0)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	l, mid, r := 0, 0, len(nums)-1

	for l <= r {
		mid = l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}

	if target < nums[mid] {
		return mid
	}

	return mid + 1
}

func searchInsert1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
