package main

import "fmt"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2023/1/13
 * Time: 23:12
 * Description: No Description
 */

func main() {
	arr := []int{2, 5, 8, 9, 11, 15, 6, 7, 7, 8, 10}
	fmt.Println(countRangeSum(arr, -1, 2))
}

func countRangeSum(nums []int, lower, upper int) int {
	if len(nums) == 0 || len(nums) < 2 {
		return 0
	}

	// 前缀和数组
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	return count(sum, 0, len(nums)-1, lower, upper)
}

func count(sum []int, l, r, lower, upper int) int {
	if l == r {
		if sum[l] >= lower && sum[l] <= upper {
			return 1
		} else {
			return 0
		}
	}

	mid := l + ((r - l) >> 2)
	leftPart := count(sum, l, mid, lower, upper)
	rightPart := count(sum, mid+1, r, lower, upper)
	merge := mergeSort(sum, l, mid, r, lower, upper)
	return leftPart + rightPart + merge
}

func mergeSort(arr []int, l, mid, r, lower, upper int) int {
	ans, windowL, windowR := 0, l, l
	for i := mid + 1; i <= r; i++ {
		min := arr[i] - upper
		max := arr[i] - lower
		for windowR <= mid && arr[windowR] <= max {
			windowR++
		}
		for windowL <= mid && arr[windowL] < min {
			windowL++
		}
		ans += maxNum(0, windowR-windowL)
	}

	help := make([]int, r-l+1)
	i, p1, p2 := 0, l, mid+1
	for p1 <= mid && p2 <= r {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			i, p1 = i+1, p1+1
		} else {
			help[i] = arr[p2]
			i, p2 = i+1, p2+1
		}
	}
	for p1 <= mid {
		help[i] = arr[p1]
		i, p1 = i+1, p1+1
	}
	for p2 <= r {
		help[i] = arr[p2]
		i, p2 = i+1, p2+1
	}
	for i := 0; i < len(help); i++ {
		arr[l+i] = help[i]
	}

	return ans
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
