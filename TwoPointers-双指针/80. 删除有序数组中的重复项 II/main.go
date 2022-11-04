package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/5
 * Time: 2:44
 * Description: No Description
 */

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
