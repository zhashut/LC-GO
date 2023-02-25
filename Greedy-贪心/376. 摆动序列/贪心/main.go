package 贪心

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/25
 * Time: 21:48
 * Description: https://leetcode.cn/problems/wiggle-subsequence/
 */

/*
https://programmercarl.com/0376.%E6%91%86%E5%8A%A8%E5%BA%8F%E5%88%97.html#%E6%80%9D%E8%B7%AF1-%E8%B4%AA%E5%BF%83%E8%A7%A3%E6%B3%95
本题异常情况的本质，就是要考虑平坡， 平坡分两种，一个是 上下中间有平坡，一个是单调有平坡
时间复杂度：O(n)
空间复杂度：O(1)
*/
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	prediff, curdiff, result := 0, 0, 1
	for i := 0; i < len(nums)-1; i++ {
		curdiff = nums[i+1] - nums[i]
		if (prediff >= 0 && curdiff < 0) || (prediff <= 0 && curdiff > 0) {
			result++
			prediff = curdiff
		}
	}

	return result
}
