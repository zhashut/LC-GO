package _5__跳跃游戏_II

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/2/28
 * Time: 20:26
 * Description: No Description
 */

// 版本一
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	curDistance, nextDistance, result := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		nextDistance = max(i+nums[i], nextDistance) // 更新下一步覆盖最远距离下标
		if i == curDistance {                       // 遇到当前覆盖最远距离下标
			if curDistance < len(nums)-1 { // 如果当前覆盖最远距离下标不是终点
				result++                         // 需要走下一步
				curDistance = nextDistance       // 更新当前覆盖最远距离下标（相当于加油了）
				if nextDistance >= len(nums)-1 { // 下一步的覆盖范围已经可以达到终点，结束循环
					break
				}
			} else { // 当前覆盖最远距到达集合终点，不用做ans++操作了，直接结束
				break
			}
		}
	}
	return result
}

// 版本二，简洁版
func jump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	curDistance, nextDistance, result := 0, 0, 0 // 当前覆盖的最远距离下标/下一步覆盖的最远距离下标/记录走的最大步数
	for i := 0; i < len(nums)-1; i++ {           // 这里是小于nums.size() - 1，这是关键所在
		nextDistance = max(i+nums[i], nextDistance) // 更新下一步覆盖的最远距离下标
		if i == curDistance {                       // 遇到当前覆盖的最远距离下标
			result++
			curDistance = nextDistance // 更新当前覆盖的最远距离下标
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
