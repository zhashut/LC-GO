package main

import "sort"

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/19
 * Time: 3:25
 * Description: No Description
 */

func main() {

}

func intersection(nums1 []int, nums2 []int) []int {
	cnt := map[int]int{}

	for _, val := range nums1 {
		cnt[val]++
	}

	ans := []int{}

	for _, val := range nums2 {
		if key, ok := cnt[val]; ok && key > 0 {
			ans = append(ans, val)
			cnt[val] = 0
		}
	}

	return ans
}

// 排序 + 双指针
func intersection1(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			y++
		}
	}

	return
}
