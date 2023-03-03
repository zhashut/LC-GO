package _35__分发糖果

/**
 * Created with GoLand 2022.2.3.
 * @author: 炸薯条
 * Date: 2023/3/3
 * Time: 23:58
 * Description: No Description
 */

func candy(ratings []int) int {
	/**先确定一边，再确定另外一边
	    1.先从左到右，当右边的大于左边的就加1
	    2.再从右到左，当左边的大于右边的就再加1
	**/
	need := make([]int, len(ratings))
	sum := 0
	// 初始化(每个人至少一个糖果)
	for i := 0; i < len(ratings); i++ {
		need[i] = 1
	}
	// 1.先从左到右，当右边的大于左边的就加1
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			need[i+1] = need[i] + 1
		}
	}
	// 2.再从右到左，当左边的大于右边的就右边加1，但要花费糖果最少，所以需要做下判断
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			need[i-1] = findMax(need[i-1], need[i]+1)
		}
	}
	//计算总共糖果
	for i := 0; i < len(ratings); i++ {
		sum += need[i]
	}
	return sum
}
func findMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
