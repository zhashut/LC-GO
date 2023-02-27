package main

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/27
 * Time: 19:52
 * Description: No Description
 */

func main() {
	nums := []int{2, 3, 1, 1, 4}
	canJump(nums)
}

func canJump(nums []int) bool {
	cover := 0
	for i := 0; i <= cover; i++ {
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
