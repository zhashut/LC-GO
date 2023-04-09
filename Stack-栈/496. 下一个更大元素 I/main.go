package _96__下一个更大元素_I

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/9
 * Time: 17:41
 * Description: https://leetcode.cn/problems/next-greater-element-i/
 * 题解：https://programmercarl.com/0496.%E4%B8%8B%E4%B8%80%E4%B8%AA%E6%9B%B4%E5%A4%A7%E5%85%83%E7%B4%A0I.html
 */

// 单调栈（未精简版本）
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range nums1 {
		res[i] = -1
	}
	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	stack := []int{0}
	for i := 0; i < len(nums2); i++ {
		top := stack[len(stack)-1]
		if nums2[i] < nums2[top] {
			stack = append(stack, i)
		} else if nums2[i] == nums2[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && nums2[i] > nums2[top] {
				if v, ok := m[nums2[top]]; ok {
					res[v] = nums2[i]
				}
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}

	return res
}

// 单调栈（精简版本）
func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}
	m := make(map[int]int, len(nums1))
	for k, v := range nums1 {
		m[v] = k
	}

	stack := []int{0}
	for i := 1; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]

			if k, ok := m[nums2[top]]; ok {
				res[k] = nums2[i]
			}

			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return res
}
