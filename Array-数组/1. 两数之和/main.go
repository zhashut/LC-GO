package main

/**
 * Created with GoLand 2021.3.2
 * @author: 炸薯条
 * Date: 2022/11/1
 * Time: 1:31
 * Description: https://leetcode.cn/problems/two-sum/
 */

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)

	for i, val := range nums {
		if p, ok := res[target-val]; ok {
			return []int{p, i}
		}
		res[val] = i
	}
	return nil
}
