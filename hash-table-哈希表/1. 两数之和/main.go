package ___两数之和

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/4/17
 * Time: 20:48
 * Description: https://leetcode.cn/problems/two-sum/
 */

func twoSum(nums []int, target int) []int {
	containerMap := make(map[int]int)

	for i, num := range nums {
		if p, ok := containerMap[target-num]; ok {
			return []int{i, p}
		}
		containerMap[num] = i
	}

	return nil
}
