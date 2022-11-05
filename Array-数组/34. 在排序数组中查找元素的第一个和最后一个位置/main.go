package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/6
 * Time: 1:22
 * Description: No Description
 */

func main() {

}

func searchRange(nums []int, target int) []int {

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res = append(res, i)
		}
	}

	if len(res) < 2 {
		res = append(res, res[0])
	}

	return res
}

// TODO: 没AC过，再改改应该也可以
func searchRange1(nums []int, target int) []int {

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	res := make([]int, 0)

	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			if nums[mid+1] == target {
				res = append(res, mid, mid+1)
			} else {
				res = append(res, mid-1, mid)
			}
			return res
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	res = append(res, -1, -1)

	return res
}

// “代码随想” 版本
func searchRange2(nums []int, target int) []int {
	leftBorder := getLeft(nums, target)
	rightBorder := getRight(nums, target)
	// 情况一
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	// 情况三
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	// 情况二
	return []int{-1, -1}
}

func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2        // 记录border没有被赋值的情况；这里不能赋值-1，target = num[0]时，会无法区分情况一和情况二
	for left <= right { // []闭区间
		mid := left + ((right - left) >> 1)
		if nums[mid] >= target { // 找到第一个等于target的位置
			right = mid - 1
			border = right
		} else {
			left = mid + 1
		}
	}
	return border
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > target {
			right = mid - 1
		} else { // 找到第一个大于target的位置
			left = mid + 1
			border = left
		}
	}
	return border

}
