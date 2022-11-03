package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/3
 * Time: 0:52
 * Description: No Description
 */

func main() {
	nums := []int{1, 1, 5, 4, 8, 7, 3}
	nextPermutation(nums)
}

func nextPermutation(nums []int) {
	i, j := 0, 0
	// 从尾部开始比较，找出小的那个点
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	// 再 [i + 1, n] 区间内找出大于 nums[i] 的数，并交换
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	//再在 [i + 1, n] 区间内找出大于 nums[i] 的数，并交换
	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
