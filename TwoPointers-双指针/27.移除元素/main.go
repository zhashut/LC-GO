package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/6
 * Time: 23:53
 * Description: https://leetcode.cn/problems/remove-element/
 */

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	removeElement1(nums, val)
}

func removeElement(nums []int, val int) int {
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if val != nums[fastIndex] {
			nums[slowIndex] = nums[fastIndex]
			slowIndex++
		}
	}
	return slowIndex
}

func removeElement1(nums []int, val int) int {
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		for leftIndex <= rightIndex && nums[leftIndex] != val {
			leftIndex++
		}
		for leftIndex <= rightIndex && nums[rightIndex] == val {
			rightIndex--
		}
		if leftIndex < rightIndex {
			nums[leftIndex] = nums[rightIndex]
			leftIndex++
			rightIndex--
		}
	}
	return leftIndex
}
