package _03__下一个更大元素_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/9
 * Time: 20:41
 * Description: https://leetcode.cn/problems/next-greater-element-ii/
 */

func nextGreaterElements(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	for i := range res {
		res[i] = -1
	}
	stack := []int{0}

	for i := 0; i < size*2; i++ {
		for len(stack) != 0 && nums[i%size] > nums[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[index] = nums[i%size]
		}
		stack = append(stack, i%size)
	}

	return res
}
